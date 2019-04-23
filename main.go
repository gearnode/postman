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
	// Specifies the date and time at which the creator of the message
	// indicated that the message was complete and ready to enter the
	// mail delivery system.  Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.1)
	Date time.Time

	// Specifies the author(s) of the message; that is, the mailbox(es)
	// of the person(s) or system(s) responsible for the writing of the
	// message.  Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.2)
	From string

	// Specifies the mailbox of the agent responsible for the actual
	// transmission of the message.  Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.2)
	Sender string

	// When the "Reply-To:" field is present, it indicates the
	// mailbox(es) to which the author of the message suggests that
	// replies be sent.  Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.2)
	ReplyTo string

	// Contains the address(es) of the primary recipient(s) of the
	// message.  Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.3)
	To []string

	// Contains the addresses of others who are to receive the message,
	// though the content of the message may not be directed at them.
	// Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.3)
	Cc []string

	// Contains addresses of recipients of the message whose addresses
	// are not to be revealed to other recipients of the message.
	// Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.3)
	Bcc []string

	// Contains a single unique message identifier that refers to a
	// particular version of a particular message.  If the message is
	// resent without changes, the original Message-ID is retained.
	// Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.4)
	MessageID string

	// The message identifier(s) of the original message(s) to which the
	// current message is a reply.  Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.4)
	InReplyTo string

	// The message identifier(s) of other message(s) to which the current
	// message may be related.  In RFC 2822, the definition was changed
	// to say that this header field contains a list of all Message-IDs
	// of messages in the preceding reply chain.  Defined as standard by
	// RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.4)
	References []string

	// Contains a short string identifying the topic of the message.
	// Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.5)
	Subject string

	// Contains any additional comments on the text of the body of the
	// message.  Warning: Some mailers will not show this field to
	// recipients.  Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s):  RFC 2822 (section 3.6.5)
	Comments []string

	// Contains a comma-separated list of important words and phrases
	// that might be useful for the recipient.  Defined as standard by
	// RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.5)
	Keywords []string

	// Contains the date and time that a message is reintroduced into the
	// message transfer system.  Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.6)
	ResentDate time.Time

	// Contains the mailbox of the agent who has reintroduced the message
	// into the message transfer system, or on whose behalf the message
	// has been resent.  Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.6)
	ResentFrom []string

	// Contains the mailbox of the agent who has reintroduced the message
	// into the message transfer system, if this is different from the
	// Resent-From value. Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.6)
	ResentSender string

	// Contains the mailbox(es) to which the message has been resent.
	// Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.6)
	ResentTo []string

	// Contains the mailbox(es) to which message is cc'ed on resend.
	// Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.6)
	ResentCc []string

	// Contains the mailbox(es) to which message is bcc'ed on resend.
	// Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.6)
	ResentBcc []string

	// Resent Reply-to. Defined by RFC 822, obsoleted by RFC 2822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822
	ResentReplyTo string

	// Contains a message identifier for a resent message. Defined as
	// standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.6)
	ResentMessageID string

	// Return path for message response diagnostics.  See also RFC 2821
	// Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.7)
	ReturnPath string

	// Contains information about receipt of the current message by a
	// mail transfer agent on the transfer path. See also RFC 2821.
	// Defined as standard by RFC 822.
	//
	// Applicable protocol: Mail
	//
	// Status: standard
	//
	// Specification document(s): RFC 2822 (section 3.6.7)
	Received string

	// Defined by RFC 822, but was found to be inadequately specified,
	// was not widely implemented, and was removed in RFC 2822.  Current
	// practice is to use separate encryption, such as S/MIME or OpenPGP,
	// possibly in conjunction with RFC 1847 MIME security multiparts.
	Encrypted string

	// Indicates that the sender wants a disposition notification when
	// this message is received (read, processed, etc.) by its
	// recipients.
	//
	// Applicable protocol: Mail
	//
	// Status: standards-track
	//
	// Specification document(s): RFC 2298
	DispositionNotificationTo string

	// For optional modifiers on disposition notification requests.
	//
	// Applicable protocol: Mail
	//
	// Status: standards-track
	//
	// Specification document(s): RFC 2298
	DispositionNotificationOptions []string

	// Indicates a language that the message sender requests to be used
	// for responses. Accept-language was not designed for email but has
	// been considered useful as input to the generation of automatic
	// replies.  Some problems have been noted concerning its use with
	// email, including but not limited to determination of the email
	// address to which it refers; cost and lack of effective
	// internationalization of email responses; interpretation of
	// language subtags; and determining what character set encoding
	// should be used.
	//
	// Applicable protocol: Mail
	//
	// Status: standards-track
	//
	// Specification document(s): RFC 3282
	AcceptLanguage string

	// A hint from the originator to the recipients about how important a
	// message is.
	//
	// Values: High, normal, or low.
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

// Output: RFC <XXX> compliant message id
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

func (m *Mail) String() (string, error) {
	var header string

	header += fmt.Sprintf("Sender: %s\r\n", m.Sender)
	header += fmt.Sprintf("From: %s\r\n", m.From)

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
		return "", err
	}

	header += fmt.Sprintf("Message-ID: %s\r\n", msgid)
	header += fmt.Sprintf("Subject: %s\r\n", m.Subject)

	return header, nil
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
