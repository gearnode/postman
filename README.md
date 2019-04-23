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

Encode in hex special chars
=> Table ref: https://www.utf8-chartable.de/unicode-utf8-table.pl

# Next

- [ ] Ensure used RFC are uptodate
- [ ] Create email package
- [ ] Create email go struct for email package
- [ ] Add Marshal func on email package (this task should be split in small steps)
- [ ] Add String on email struct as Marshal alias func

# References
- https://tools.ietf.org/html/rfc4021#section-1
- https://tools.ietf.org/html/rfc6532
- https://tools.ietf.org/html/rfc5322
- https://www.ietf.org/rfc/rfc2045.txt
- https://www.ietf.org/rfc/rfc2046.txt
- https://www.ietf.org/rfc/rfc2047.txt
- https://www.ietf.org/rfc/rfc2048.txt
- https://www.ietf.org/rfc/rfc2049.txt

# Requirements

A fake SMTP server.

```sh
docker run -d -p 1080:1080 -p 1025:1025 --name mailcatcher2 schickling/mailcatcher
```
