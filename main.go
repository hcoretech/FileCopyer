package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type file struct {
	byteOfData []byte
}

func getByte(val *file) *file {
	return val
}
func setByte(val []byte) *file {
	byteFirstData := &file{val}
	return byteFirstData
}

func (val *file) addByte(newByte []byte) *file {
	if val.byteOfData != nil {
		nByte := newByte
		base := val
		base.byteOfData = append(base.byteOfData, nByte...)
		fmt.Println("adding data to cache")
		return base
	}

	d := setByte(newByte)
	fmt.Println("setting data for the first time")
	return d

}

func main() {
	fmt.Println("STARTING FILE COPYER")
	arg := os.Args[1]
	totalSize := make(chan int)
	done := make(chan bool)
	fmt.Println("Reading Passed Arquement from terminal")
	if arg != "" {
		path := filepath.Join("C:\\Users\\HCORE\\Downloads\\", arg)
		path2 := filepath.Join("C:\\Users\\HCORE\\Desktop", arg)
		time.Sleep(2 * time.Second)

		fmt.Println("opening the file ")
		open, err := os.Open(path)
		if err != nil {
			log.Fatal("error opening file")
		}
		defer open.Close()
		time.Sleep(1 * time.Second)
		fmt.Println("Getting file info")
		fileInfo, err := open.Stat()
		if err != nil {
			log.Fatal("error getting file info", err)
			return
		}
		time.Sleep(1 * time.Second)
		getSize := int(fileInfo.Size())
		fmt.Println("file path", open.Name())
		fmt.Println("file opened successfully", getSize)
		time.Sleep(1 * time.Second)
		fmt.Println("Reading File Data")
		go func() {
			for {
				size, ok := <-totalSize
				if ok {
					fmt.Println("recieve data of size", size)
				} else {
					fmt.Println("done recieving data", getSize)
					done <- true
					return
				}
			}

		}()

		Percent := getSize * 10 / 1024 * 256
		for i := 0; i <= getSize; i += Percent {
			time.Sleep(time.Second)
			in64 := int64(i)
			_, err := open.Seek(in64, Percent)
			if err != nil {
				log.Fatal("error seeking file", err)
			}
			fmt.Println("sending  data", i)
			totalSize <- i

		}

		close(totalSize)
		<-done
		fmt.Println("file read and stored completely")
		time.Sleep(1 * time.Second)
		fmt.Println("creating new file path")
		time.Sleep(1 * time.Second)
		newFile, err := os.Create(path2)
		if err != nil {
			fmt.Println("error creating file", err)
			return
		}
		defer newFile.Close()
		time.Sleep(2 * time.Second)
		fmt.Println("file created successfully in the desktop directory  ", open.Name())
		time.Sleep(2 * time.Second)
		fmt.Println("copying byte")
		err = os.Chmod(path2, fileInfo.Mode().Perm())
		if err != nil {
			log.Fatal("failed to set destination file permissions: %w", err)
		}
		bytesCopied, err := io.Copy(newFile, open)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Copied %d bytes\n", bytesCopied)

		// // 6. Sync the destination file to disk
		err = newFile.Sync()
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(2 * time.Second)
		fmt.Println("file copied successfully")

	}

}

// func main ( ) {
//     fmt.Println("*^*")
//    fmt.Println("^^^^^^")
//  fmt.Println("^^^^^^^^^^")
// fmt.Println("^^^^^^^^^^^^^")
//   fmt.Println("^^^^^^^^^^")
// 	 fmt.Println("^^^^")

// }

// func main () {
// 	interval:= time.NewTicker(1000 * time.Millisecond)
// 	interval2 := time.NewTicker(10000 * time.Millisecond)
// 	done := make(chan bool)
// 	isDone := make(chan bool)

// 	go func() {
// 		for  {
// 		select{
// 			case <- done:
// 				return
// 			case t := <- interval.C:
// 				fmt.Println("Ticker start at ", t)
// 		}
// 		}
// 	}()
// 	go func() {
// 		for  {
// 		select{
// 			case  dones:= <- isDone:
// 					fmt.Println("Interval channel closed", dones )
// 						return
// 				//  }
// 			case t := <- interval2.C:
// 				fmt.Println("Tickers closed at " ,t)

// 		}
// 		}
// 	}()
// 	time.Sleep(10000 * time.Millisecond)
// 	interval.Stop()
// 	done <- true

//     time.Sleep( time.Millisecond)
// 	interval2.Stop()
// 	isDone <- true

// 	// time.Sleep(2 * time.Second)
// }

// func main () {
// 	timer1 := time.NewTimer(2 * time.Second)
// 	<- timer1.C
// 	fmt.Println("Timer 1 expired")
// 	timer2 := time.NewTimer( time.Second)

// 		// fmt.Println("Timer 2 expired")

// 	go func() {
// 		<- timer2.C
// 		fmt.Println("Timer 2 expired")
// 	}()
// 	stop := timer2.Stop()
// 	//  fmt.Println("Timer 2 stopped")
// 	if stop {
// 		fmt.Println("Timer 2 stopped")
// 	}
// 	//  time.Sleep(2 * time.Second)

// }

// func main () {
// 	fmt.Println("HCORE MESSAGING SYSTEM")
// 	message := make(chan string)
// 	messageInt := make (chan int, 5)
// 	isSent := make(chan bool)

// 	go func(){
// 		for{
// 			msgInt , ok := <- messageInt
// 			msg , _ := <-message
// 			if ok {
// 			   fmt.Printf("recieved %v %v\n ",msg,msgInt)
// 			}else{
// 				fmt.Println("message sent complete")
// 				isSent <- true
// 				return
// 			}
// 		}

// 	}()

// 	for i := 1 ;i <= 4 ; i++{
// 		time.Sleep(3 * time.Second)
// 	  fmt.Printf("sent henry message %v\n",i)
// 	   messageInt <-i
//        message <- "henry message"
// 	   time.Sleep(4*time.Second)
// 	}
// 	close(message)
// 	close(messageInt)

//     // time.Sleep( * time.Second)
//     <-isSent

// 	_,ok := <-messageInt
// 	fmt.Println("message exiting" , ok)

// }
