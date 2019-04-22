package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type Mail struct {
	//  Contains a single unique message identifier that refers to a
	// particular version of a particular message.  If the message is
	// resent without changes, the original Message-ID is retained.
	// Defined as standard by RFC 822.
	MessageID string

	// Specifies the author(s) of the message; that is, the mailbox(es)
	// of the person(s) or system(s) responsible for the writing of the
	// message.  Defined as standard by RFC 822.
	From string

	// Specifies the mailbox of the agent responsible for the actual
	// transmission of the message.  Defined as standard by RFC 822.
	Sender string

	// Contains the address(es) of the primary recipient(s) of the
	// message.  Defined as standard by RFC 822.
	To []string

	// Contains the addresses of others who are to receive the message,
	// though the content of the message may not be directed at them.
	// Defined as standard by RFC 822.
	Cc []string

	// Contains addresses of recipients of the message whose addresses
	// are not to be revealed to other recipients of the message.
	// Defined as standard by RFC 822.
	Bcc []string

	// Contains a short string identifying the topic of the message.
	// Defined as standard by RFC 822.
	Subject string

	// When the "Reply-To:" field is present, it indicates the
	// mailbox(es) to which the author of the message suggests that
	// replies be sent.  Defined as standard by RFC 822.
	ReplyTo string

	// The message identifier(s) of the original message(s) to which the
	// current message is a reply.  Defined as standard by RFC 822.
	InReplyTo string

	// The message identifier(s) of other message(s) to which the current
	// message may be related.  In RFC 2822, the definition was changed
	// to say that this header field contains a list of all Message-IDs
	// of messages in the preceding reply chain.  Defined as standard by
	// RFC 822.
	References []string

	// Contains any additional comments on the text of the body of the
	// message.  Warning: Some mailers will not show this field to
	// recipients.  Defined as standard by RFC 822.
	Comments string

	// Contains a comma-separated list of important words and phrases
	// that might be useful for the recipient.  Defined as standard by
	// RFC 822.
	Keywords []string

	// A hint from the originator to the recipients about how important a
	// message is.
	//
	// Values: High, normal, or low.
	//
	// Not used to control transmission speed. Proposed for use with RFC
	// 2156 (MIXER) [10] and RFC 3801 (VPIM) [14].
	Importance string

	// Can be 'normal', 'urgent', or 'non-urgent' and can influence
	// transmission speed and delivery.  RFC 2156 (MIXER); not for
	// general use.
	Priority string

	// How sensitive it is to disclose this message to people other than
	// the specified recipients.  Values: Personal, private, and company
	// confidential.  The absence of this header field in messages
	// gatewayed from X.400 indicates that the message is not sensitive.
	// Proposed for use with RFC 2156 (MIXER) [10] and RFC 3801 (VPIM)
	// [14].
	Sensitivity string

	Parts []Part

	Attachments []Attachment
}

type Part struct {
	ContentType string

	Content []byte
}

type Attachment struct {
	Filename string

	ContentDisposition string

	ContentID string

	ContentTransfertEncoding string

	Content []byte
}

func genMsgID() (string, error) {
	t := time.Now().UnixNano()
	pid := os.Getpid()

	// TODO: @geanode what's the cost of NewInt with big package
	rint, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return "", err
	}

	host, err := os.Hostname()
	if err != nil {
		host = "localhost.localdomain"
	}

	return fmt.Sprintf("<%d.%d.%d@%s>", t, pid, rint, host), nil
}

func (m *Mail) Marshal() {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", m.Sender)

	if len(m.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(m.To, ";"))
	}

	if len(m.Cc) > 0 {
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(m.Cc, ";"))
	}

	if len(m.Bcc) > 0 {
		header += fmt.Sprintf("Bcc: %s\r\n", strings.Join(m.Bcc, ";"))
	}

	if m.ReplyTo != "" {
		header += fmt.Sprintf("Reply-To: %s\r\n", m.ReplyTo)
	}

	msgid, err := genMsgID()
	if err != nil {
		// TODO: @gearnode properly handle error
		panic(err)
	}

	header += fmt.Sprintf("Message-ID: %s\r\n", msgid)

	header += fmt.Sprintf("Subject: %s\r\n", m.Subject)
}

func main() {
	conn, err := smtp.Dial("localhost:1025")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	conn.Mail("foo@bar.fr")
	conn.Rcpt("recp@foo.fr")

	wc, err := conn.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()

	buf := bytes.NewBufferString("some email body")

	_, err = buf.WriteTo(wc)
	if err != nil {
		log.Fatal(err)
	}
}
