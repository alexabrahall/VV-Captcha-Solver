import { solveVVCaptcha } from "../src/typescript/index";

async function main() {
  try {
    const token = await solveVVCaptcha(
      "https://uk-umg.com/um-forms/48709-1272284.html"
    );
    console.log(`VV Captcha token: ${token}`);
  } catch (error) {
    console.error(error);
  }
}
