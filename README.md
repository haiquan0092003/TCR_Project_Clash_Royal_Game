# 🎮 Game Server-Client TCP

<div style="background-color: #f8f9fa; padding: 20px; border-radius: 10px; margin: 20px 0;">
    <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 10px;">Tính Năng Chính</h2>
    <ul style="list-style-type: none; padding-left: 0;">
        <li style="margin: 10px 0; padding: 10px; background: white; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
            <span style="color: #e74c3c;">🎯</span> Hệ thống đăng nhập và xác thực
        </li>
        <li style="margin: 10px 0; padding: 10px; background: white; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
            <span style="color: #e74c3c;">🎮</span> Quản lý phiên chơi game
        </li>
        <li style="margin: 10px 0; padding: 10px; background: white; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
            <span style="color: #e74c3c;">📊</span> Hệ thống xếp hạng
        </li>
        <li style="margin: 10px 0; padding: 10px; background: white; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
            <span style="color: #e74c3c;">🔒</span> Bảo mật và mã hóa
        </li>
    </ul>
</div>

<div style="background-color: #f8f9fa; padding: 20px; border-radius: 10px; margin: 20px 0;">
    <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 10px;">Cài Đặt</h2>
    <div style="background: white; padding: 15px; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <h3 style="color: #2c3e50;">Yêu Cầu Hệ Thống</h3>
        <ul>
            <li>Go 1.16 trở lên</li>
            <li>Windows/Linux/MacOS</li>
        </ul>
    </div>
    
    <div style="background: white; padding: 15px; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); margin-top: 15px;">
        <h3 style="color: #2c3e50;">Cài Đặt Dependencies</h3>
        <pre style="background: #f1f1f1; padding: 10px; border-radius: 5px;"><code>go mod download</code></pre>
    </div>
    
    <div style="background: white; padding: 15px; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); margin-top: 15px;">
        <h3 style="color: #2c3e50;">Biên Dịch</h3>
        <pre style="background: #f1f1f1; padding: 10px; border-radius: 5px;"><code>go build -o server.exe server.go
go build -o client.exe client.go</code></pre>
    </div>
</div>

<div style="background-color: #f8f9fa; padding: 20px; border-radius: 10px; margin: 20px 0;">
    <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 10px;">Cách Chơi</h2>
    <div style="background: white; padding: 15px; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <h3 style="color: #2c3e50;">Khởi Động Server</h3>
        <pre style="background: #f1f1f1; padding: 10px; border-radius: 5px;"><code>./server.exe</code></pre>
    </div>
    
    <div style="background: white; padding: 15px; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); margin-top: 15px;">
        <h3 style="color: #2c3e50;">Khởi Động Client</h3>
        <pre style="background: #f1f1f1; padding: 10px; border-radius: 5px;"><code>./client.exe</code></pre>
    </div>
    
    <div style="background: white; padding: 15px; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); margin-top: 15px;">
        <h3 style="color: #2c3e50;">Các Lệnh Trong Game</h3>
        <ul>
            <li><code>/login [username] [password]</code> - Đăng nhập</li>
            <li><code>/register [username] [password]</code> - Đăng ký</li>
            <li><code>/play</code> - Bắt đầu chơi game</li>
            <li><code>/quit</code> - Thoát game</li>
        </ul>
    </div>
</div>

<div style="background-color: #f8f9fa; padding: 20px; border-radius: 10px; margin: 20px 0;">
    <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 10px;">Cấu Trúc Dự Án</h2>
    <div style="background: white; padding: 15px; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <pre style="background: #f1f1f1; padding: 10px; border-radius: 5px;"><code>.
├── server.go          # Server chính
├── client.go          # Client chính
├── clients.go         # Quản lý client
├── json_utils.go      # Xử lý JSON
├── level_utils.go     # Quản lý level
├── go.mod            # Quản lý dependencies
└── README.md         # Tài liệu</code></pre>
    </div>
</div>

<div style="background-color: #f8f9fa; padding: 20px; border-radius: 10px; margin: 20px 0;">
    <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 10px;">Bảo Mật</h2>
    <div style="background: white; padding: 15px; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <ul>
            <li>Mã hóa dữ liệu truyền tải</li>
            <li>Xác thực người dùng</li>
            <li>Bảo vệ chống tấn công</li>
        </ul>
    </div>
</div>

<div style="background-color: #f8f9fa; padding: 20px; border-radius: 10px; margin: 20px 0;">
    <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 10px;">Đóng Góp</h2>
    <div style="background: white; padding: 15px; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <p>Mọi đóng góp đều được hoan nghênh! Vui lòng tạo issue hoặc pull request.</p>
    </div>
</div>

<div style="background-color: #f8f9fa; padding: 20px; border-radius: 10px; margin: 20px 0;">
    <h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 10px;">Giấy Phép</h2>
    <div style="background: white; padding: 15px; border-radius: 5px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <p>MIT License</p>
    </div>
</div>

<style>
    body {
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        line-height: 1.6;
        color: #333;
    }
    
    h1 {
        color: #2c3e50;
        text-align: center;
        font-size: 2.5em;
        margin-bottom: 30px;
    }
    
    h2 {
        font-size: 1.8em;
        margin-top: 30px;
    }
    
    h3 {
        font-size: 1.4em;
        margin-top: 20px;
    }
    
    code {
        background-color: #f1f1f1;
        padding: 2px 5px;
        border-radius: 3px;
        font-family: 'Courier New', Courier, monospace;
    }
    
    pre {
        margin: 15px 0;
        overflow-x: auto;
    }
    
    ul {
        padding-left: 20px;
    }
    
    li {
        margin: 5px 0;
    }
</style> 