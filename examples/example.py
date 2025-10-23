import sys
import os
sys.path.append(os.path.join(os.path.dirname(__file__), '..'))

from src.python.main import solve_vv_captcha

def main():
    try:
        token = solve_vv_captcha("https://uk-umg.com/um-forms/48709-1272284.html")
        print(f"VVCaptcha token: {token}")
    except Exception as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    main()