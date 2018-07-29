package parser

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func ParseURL(input *http.Response) string {
	if input == nil {
		return ""
	}
	b, err := ioutil.ReadAll(input.Body)
	defer input.Body.Close()
	if err != nil {
		return ""
	}
	html := string(b[:])
	re := regexp.MustCompile(`http(?:s)?\:\/\/stackoverflow\.com.*?\"`)
	matches := re.FindString(html)
	if len(matches) <= 0 {
		return ""
	}
	return matches[:len(matches)-1]
}

func ParseAnswer(input *http.Response) string {
	if input == nil {
		return ""
	}
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

	reAnswer = regexp.MustCompile(`(?s)\<code(\s+.*?)?\>(.*?)\<\/code\>`)
	answer = reAnswer.ReplaceAllString(answer, "<yellow>$2</yellow>")

	reAnswer = regexp.MustCompile(`\<strong(\s+.*?)?\>(.*?)\<\/strong\>`)
	answer = reAnswer.ReplaceAllString(answer, "<cyan>$2</cyan>")

	reAnswer = regexp.MustCompile(`(?s)\<em(\s+.*?)?\>(.*?)\<\/em\>`)
	answer = reAnswer.ReplaceAllString(answer, "<cyan>$2</cyan>")

	reAnswer = regexp.MustCompile(`(?s)\<sup(\s+.*?)?\>(.*?)\<\/sup\>`)
	answer = reAnswer.ReplaceAllString(answer, "<cyan>$2</cyan>")

	reAnswer = regexp.MustCompile(`(?s)\<(h1|h2|h3|h4|h5|h6)(\s+.*?)?\>(.*?)\<\/(h1|h2|h3|h4|h5|h6)\>`)
	answer = reAnswer.ReplaceAllString(answer, "<cyan><u>$3</u></cyan>")

	reAnswer = regexp.MustCompile(`(?s)\<blockquote(\s+.*?)?\>(.*?)\<\/blockquote\>`)
	answer = reAnswer.ReplaceAllString(answer, "<green>$2</green>")

	reAnswer = regexp.MustCompile(`(?s)\<p(\s+.*?)?\>(.*?)\<\/p\>`)
	answer = reAnswer.ReplaceAllString(answer, "$2")

	reAnswer = regexp.MustCompile(`(?s)\<pre(\s+.*?)?\>(.*?)\<\/pre\>`)
	answer = reAnswer.ReplaceAllString(answer, "$2")

	reAnswer = regexp.MustCompile(`&lt;`)
	answer = reAnswer.ReplaceAllString(answer, "<")

	reAnswer = regexp.MustCompile(`&gt;`)
	answer = reAnswer.ReplaceAllString(answer, ">")

	reAnswer = regexp.MustCompile(`\<a(\s+.*?)?href=\"(.*?)\"\>.*?\<\/a\>`)
	answer = reAnswer.ReplaceAllString(answer, "<blue>$2</blue>")

	reAnswer = regexp.MustCompile(`\".rel=\"nofollow`)
	answer = reAnswer.ReplaceAllString(answer, "")

	reAnswer = regexp.MustCompile(`(?s)\<(ul|ol)(\s+.*?)?\>(.*?)\<\/(ul|ol)\>`)
	answer = reAnswer.ReplaceAllString(answer, "$3")

	reAnswer = regexp.MustCompile(`(?s)\<li(\s+.*?)?\>(.*?)\<\/li\>`)
	answer = reAnswer.ReplaceAllString(answer, "  <green>-</green> $2")

	reAnswer = regexp.MustCompile(`\<\/?hr\/?\>`)
	answer = reAnswer.ReplaceAllString(answer, "------------------------------------------------------------")

	output := "<green><u>Question:</u></green> " + question + "\n\n"
	output += "<green><u>Answer:</u></green>\n" + answer

	return output
}
