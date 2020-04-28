package pagefile

func (file *File) run() {
	defer file.wg.Done()

	for {
		select {
		case <-file.exitSig:
			return
		case line := <-file.lineCh:
			file.fd.WriteString(line + "\n")
			file.page.lineCnt++

			if file.page.isFull() {
				file.page = file.page.next()
				file.openOrCreate()
			}
		}
	}
}

func (file *File) Close() {
	close(file.exitSig)
	file.wg.Wait()
	file.fd.Close()
}
