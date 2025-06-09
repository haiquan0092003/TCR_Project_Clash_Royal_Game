// Package main chứa code cho client game
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// Các hằng số cho việc kết nối lại
const (
	maxRetries = 3               // Số lần thử kết nối tối đa
	retryDelay = 2 * time.Second // Thời gian chờ giữa các lần thử
)

// connectToServer thử kết nối đến server với số lần thử lại có giới hạn
func connectToServer() (net.Conn, error) {
	var conn net.Conn
	var err error

	// Thử kết nối tối đa maxRetries lần
	for i := 0; i < maxRetries; i++ {
		conn, err = net.Dial("tcp", "localhost:8080")
		if err == nil {
			return conn, nil
		}
		fmt.Printf("Lần thử kết nối %d thất bại: %v\n", i+1, err)
		if i < maxRetries-1 {
			fmt.Printf("Đang thử kết nối lại sau %v...\n", retryDelay)
			time.Sleep(retryDelay)
		}
	}
	return nil, fmt.Errorf("không thể kết nối sau %d lần thử", maxRetries)
}

// handleConnection xử lý kết nối với server và tương tác người dùng
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	stdin := bufio.NewReader(os.Stdin)

	// Vòng lặp xử lý đăng nhập
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Mất kết nối trong quá trình đăng nhập:", err)
			return
		}
		fmt.Print(msg)

		// Chỉ đọc input khi server yêu cầu (chứa từ "Nhập")
		if strings.Contains(msg, "Nhập") {
			input, _ := stdin.ReadString('\n')
			fmt.Fprint(conn, input)
		}

		// Thoát vòng lặp khi đăng nhập thành công hoặc thất bại
		if strings.Contains(msg, "Sai thông tin") || strings.Contains(msg, "thành công") {
			break
		}
	}

	// Vòng lặp xử lý game
	for {
		// Set timeout cho mỗi lần đọc để tránh treo
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))

		msg, err := reader.ReadString('\n')
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				// Nếu timeout, kiểm tra kết nối còn sống không
				if _, err := conn.Write([]byte{}); err != nil {
					fmt.Println("Mất kết nối với server:", err)
					return
				}
				continue
			}
			fmt.Println("Mất kết nối trong game:", err)
			return
		}
		fmt.Print(msg)

		// Xử lý input khi server yêu cầu (chứa từ "Chọn")
		if strings.Contains(msg, "Chọn") {
			input, _ := stdin.ReadString('\n')
			_, err := fmt.Fprint(conn, input)
			if err != nil {
				fmt.Println("Lỗi gửi input:", err)
				return
			}
		}

		// Kết thúc game khi nhận thông báo kết thúc
		if strings.Contains(msg, "Trận đấu kết thúc") {
			// Đợi thêm tin nhắn cuối cùng từ server
			time.Sleep(1 * time.Second)
			return
		}
	}
}

// main là điểm khởi đầu của chương trình
func main() {
	for {
		// Thử kết nối đến server
		conn, err := connectToServer()
		if err != nil {
			fmt.Println("Không thể kết nối tới server:", err)
			return
		}

		// Xử lý kết nối và chơi game
		handleConnection(conn)

		// Hỏi người chơi có muốn chơi lại không
		fmt.Print("\nBạn có muốn chơi lại không? (y/n): ")
		var choice string
		fmt.Scanln(&choice)
		if strings.ToLower(choice) != "y" {
			break
		}
		fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	}
}
