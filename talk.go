// Package talk prints an ASCII image with a message.
package talk

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// Gopher writes an ASCII image of a gopher with a message.
func Gopher(s string, w io.Writer) {
	ASCII := []string{
		"                        dhyysoo++++ooossyhd",
		"                Ndyo/:#######################:+sd",
		"      NN/:###ds/#######################/+////+/###+y#######",
		"  Ny/:###/hy:#:+/::###::/+:#########+/:`  .:#  :+/##:oo######",
		" h:##//:+o:#:+:  +hdy.    #+:#####++`     s####` :+###:sN +###:",
		" ###d   /##:o   :#####      o:###//       h#s#N.  :/####+Ny####",
		"d###sN :###s`    `sd+d:     `o###s         ://`    s#####++###/",
		" y:##s/####s                `s###s`               `s######s:/",
		"  Nhso#####//               +:###:o`             `o:######:",
		"    N#######+/            `+/:+sss+o/`         ./+#########",
		"    y########:+/:.    `#://#+N     N:+//:::::/+/###########/",
		"    +############:/+++/:##ooodNNN dsoo+#####################",
		"    /####################s/:::::::::::/s####################h",
		"    :####################+o//+++s/++//oo####################s",
		"    :######################o/  ##  ++/:#####################o",
		"    /######################:+  ::  :/#######################+",
		"    +#######################+//++//+########################+",
	}

	if s == "" {
		return
	}

	io.WriteString(w, talk(s))
	io.WriteString(w, strings.Join(ASCII, "\n"))
	io.WriteString(w, "\n")
}

func talk(s string) string {
	var buf bytes.Buffer
	var maxLen int

	in := bufio.NewScanner(strings.NewReader(s))
	dialog := []string{}

	for in.Scan() {
		line := in.Text()

		if len(line) > maxLen {
			maxLen = len(line)
		}

		dialog = append(dialog, line)
	}

	buf.WriteString(" " + strings.Repeat("~", maxLen+3) + "\n")

	for _, line := range dialog {
		buf.WriteString("/ " + line + strings.Repeat(" ", maxLen-len(line)) + " /" + "\n")
	}

	buf.WriteString(strings.Repeat("~", maxLen+3) + "\n")

	for i := 3; i > 0; i-- {
		buf.WriteString(strings.Repeat(" ", maxLen) + strings.Repeat("#", i) + "\n")
	}

	return buf.String()
}
