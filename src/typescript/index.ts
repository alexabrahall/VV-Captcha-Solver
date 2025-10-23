import axios from "axios";

export async function solveVVCaptcha(url: string) {
  const domain = new URL(url).hostname; // grab domain for base url

  const response = await axios.get(
    `https://${domain}/handler/vv_captcha/code`,
    {
      headers: {
        "User-Agent":
          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36",
        Referer: url,
      },
      validateStatus(status) {
        return status >= 200 && status < 500;
      },
    }
  );

  //status code handling

  if (response.status === 403) {
    //error in url
    throw new Error(
      "Invalid URL - Domain verification failed (this domain is likely not a vv captcha enabled domain)"
    );
  } else if (response.status === 500) {
    //failed to verify captcha
    throw new Error(
      "Failed to verify captcha - likely error in domain handling"
    );
  } else if (response.status !== 200) {
    // unexpected response
    throw new Error(
      `Unexpected error [status: ${response.status}] - Please check your domain settings`
    );
  }

  const captchaData = response.data;

  if (!captchaData) {
    throw new Error(
      "Failed to fetch captcha data - likely error in domain handling"
    );
  }

  //regex pattern to extract the captcha id from the response

  const captchaIdMatch = captchaData.match(
    /var\s+token\s*=\s*(['"])(.*?)\1\s*;/
  );

  if (!captchaIdMatch || captchaIdMatch.length < 3) {
    throw new Error("Failed to extract captcha id from response");
  }

  const captchaId = captchaIdMatch[2];

  return captchaId;
}
