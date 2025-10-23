import requests
import re
from urllib.parse import urlparse

def solve_vv_captcha(url: str) -> str:
    domain = urlparse(url).hostname  # grab domain for base url
    
    response = requests.get(
        f"https://{domain}/handler/vv_captcha/code",
        headers={
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36",
            "Referer": url
        }
    )
    
    # status code handling
    if response.status_code == 403:
        # error in url
        raise Exception("Invalid URL - Domain verification failed (this domain is likely not a vv captcha enabled domain)")
    elif response.status_code == 500:
        # failed to verify captcha
        raise Exception("Failed to verify captcha - likely error in domain handling")
    elif response.status_code != 200:
        # unexpected response
        raise Exception(f"Unexpected error [status: {response.status_code}] - Please check your domain settings")
    
    captcha_data = response.text
    
    if not captcha_data:
        raise Exception("Failed to fetch captcha data - likely error in domain handling")
    
    # regex pattern to extract the captcha id from the response
    captcha_id_match = re.search(r'var\s+token\s*=\s*([\'"])(.*?)\1\s*;', captcha_data)
    
    if not captcha_id_match or len(captcha_id_match.groups()) < 2:
        raise Exception("Failed to extract captcha id from response")
    
    captcha_id = captcha_id_match.group(2)
    
    return captcha_id

