# Architecture & Design
![alt text](https://raw.githubusercontent.com/chapterzero/tn_test/master/doc/architecture.jpg)

User access the client using his browser, and the browser will make call to API, which contain logic, validation and database operation. Client & API is served by main program `server.go`.

Notification job should go to queue server first not directly to the sender service (SMTP / push notification service etc) to prevent issue like long process time when creating transaction because smtp server not able to contact user email host and complication to resend failed job - where using queue server, requeue is handled by the queue server. Also this architecture allow more flexible for future notification development - for example if you want to add sms, only need to rework on the job consumer, not every existing endpoint.

# Setup
- Require mysql server, or use docker-compose environment using `docker-compose up -d`
- Edit app.config.json with your database credential
- Initialize your database with `sql/init.sql` file.
- Run the server with `go run server.go`
- Go to your browser, type `http://localhost:8777/`
