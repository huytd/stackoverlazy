package parser

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func ParseURL(input *http.Response) string {
	b, err := ioutil.ReadAll(input.Body)
	defer input.Body.Close()
	if err != nil {
		return ""
	}
	html := string(b[:])
	re := regexp.MustCompile(`http\:\/\/stackoverflow\.com.*?\"`)
	matches := re.FindString(html)
	if len(matches) <= 0 {
		return ""
	}
	return matches[:len(matches)-1]
}

func ParseAnswer(input *http.Response) string {
	b, err := ioutil.ReadAll(input.Body)
	defer input.Body.Close()
	if err != nil {
		return ""
	}
	html := string(b[:])

	reQuestion := regexp.MustCompile(`class=\"question-hyperlink\"\>(.*?)\<\/a\>`)
	question := ""
	questionArray := reQuestion.FindStringSubmatch(html)
	if len(questionArray) > 1 {
		question = questionArray[1]
	}

	reAnswer := regexp.MustCompile(`(?s)class=\"answercell\"\>.*?itemprop=\"text\"\>(.*?)\<\/div\>`)
	answer := ""
	answerArray := reAnswer.FindStringSubmatch(html)
	if len(answerArray) > 1 {
		answer = answerArray[1]
	}

	reAnswer = regexp.MustCompile(`(?s)\<code\>(.*?)\<\/code\>`)
	answer = reAnswer.ReplaceAllString(answer, "<yellow>$1</yellow>")

	reAnswer = regexp.MustCompile(`\<strong\>(.*?)\<\/strong\>`)
	answer = reAnswer.ReplaceAllString(answer, "<cyan>$1</cyan>")

	reAnswer = regexp.MustCompile(`(?s)\<em\>(.*?)\<\/em\>`)
	answer = reAnswer.ReplaceAllString(answer, "<cyan>$1</cyan>")

	reAnswer = regexp.MustCompile(`(?s)\<sup\>(.*?)\<\/sup\>`)
	answer = reAnswer.ReplaceAllString(answer, "<cyan>$1</cyan>")

	reAnswer = regexp.MustCompile(`(?s)\<(h1|h2|h3|h4|h5|h6)\>(.*?)\<\/(h1|h2|h3|h4|h5|h6)\>`)
	answer = reAnswer.ReplaceAllString(answer, "<cyan><u>$2</u></cyan>")

	reAnswer = regexp.MustCompile(`(?s)\<blockquote\>(.*?)\<\/blockquote\>`)
	answer = reAnswer.ReplaceAllString(answer, "<green>$1</green>")

	reAnswer = regexp.MustCompile(`(?s)\<p\>(.*?)\<\/p\>`)
	answer = reAnswer.ReplaceAllString(answer, "$1")

	reAnswer = regexp.MustCompile(`(?s)\<pre\>(.*?)\<\/pre\>`)
	answer = reAnswer.ReplaceAllString(answer, "$1")

	reAnswer = regexp.MustCompile(`&lt;`)
	answer = reAnswer.ReplaceAllString(answer, "<")

	reAnswer = regexp.MustCompile(`&gt;`)
	answer = reAnswer.ReplaceAllString(answer, ">")

	reAnswer = regexp.MustCompile(`\<a href=\"(.*?)\"\>.*?\<\/a\>`)
	answer = reAnswer.ReplaceAllString(answer, "<blue>$1</blue>")

	reAnswer = regexp.MustCompile(`\".rel=\"nofollow`)
	answer = reAnswer.ReplaceAllString(answer, "")

	reAnswer = regexp.MustCompile(`(?s)\<(ul|ol)\>(.*?)\<\/(ul|ol)\>`)
	answer = reAnswer.ReplaceAllString(answer, "$2")

	reAnswer = regexp.MustCompile(`(?s)\<li\>(.*?)\<\/li\>`)
	answer = reAnswer.ReplaceAllString(answer, "  <green>-</green> $1")

	reAnswer = regexp.MustCompile(`\<\/?hr\/?\>`)
	answer = reAnswer.ReplaceAllString(answer, "------------------------------------------------------------")

	output := "<green><u>Question:</u></green> " + question + "\n\n"
	output += "<green><u>Answer:</u></green>\n" + answer

	return output
}
