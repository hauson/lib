package pagefile

func (file *PageFile) run() {
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

func (file *PageFile) Close() {
	close(file.exitSig)
	file.wg.Wait()
	file.fd.Close()
}
