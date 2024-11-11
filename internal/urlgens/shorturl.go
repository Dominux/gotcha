package urlgens

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func GenShortUrl(link string) (string, error) {
	// making request
	var resp *http.Response
	{
		reqBody := new(bytes.Buffer)
		mp := multipart.NewWriter(reqBody)
		mp.WriteField("u", link)
		mp.Close()

		req, err := http.NewRequest(http.MethodPost, "https://www.shorturl.at/shortener.php", reqBody)
		if err != nil {
			return "", errors.New("WTF?")
		}
		req.Header.Set("Content-Type", mp.FormDataContentType())

		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			return "", errors.New("[shorturl generator API error]" + err.Error())
		}
	}

	// reading response body
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyHtml := string(bodyBytes)

	// parsing html
	parts := strings.Split(bodyHtml, `<input id="shortenurl" type="text" value="`)
	if len(parts) != 2 {
		println(bodyHtml)
		return "", errors.New("shorturl returned wrong html")
	}

	parts = strings.Split(parts[1], `"`)
	if len(parts) == 1 {
		return "", errors.New("shorturl returned whoever knows wht")
	}

	return parts[0], nil
}
