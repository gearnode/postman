# Status

For now this project is only a playground to play with SMTP and golang.

# Note

Golang does not have simple and maintain library to marshal and unmarshal
email on SMTP socket.

=> Write a basic marshal function to send email easy with Golang

The net/smtp package does not allow to send many email with the same TCP
socket.

=? Hack the std to send many email on the same socket
=? Search an package to send many email with the same socket
=? Write a SMTP protocol package
