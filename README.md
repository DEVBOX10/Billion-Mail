<div align="center">
  <a name="readme-top"></a>
  <h1><a href="https://www.billionmail.com/" target="_blank">BillionMail 📧</a></h1>

## An Open-Source Mail Server, Email Marketing Solution for Smarter Campaigns

[![][license-shield]][license-link]
[![][docs-shield]][docs-link]
[![][github-release-shield]][github-release-link]
[![][github-stars-shield]][github-stars-link]

</div>
<br/>

## What is BillionMail?

BillionMail is a **future open-source Mail server, Email marketing platform** designed to help businesses and individuals manage their email campaigns with ease. Whether you're sending newsletters, promotional emails, or transactional messages, this tool will provide **full control** over your email marketing efforts. With features like **advanced analytics**, and **customer management**, you'll be able to create, send, and track emails like a pro.

![Billion Mail Banner](https://www.billionmail.com/home.png?v1)

## How to use?
**Install Script:** (✅The script automatically installs all required runtime environments including Docker)
```shell
cd /opt
git clone https://github.com/aaPanel/Billion-Mail
cd Billion-Mail

# install with interact
bash install.sh
```


**Install with Docker:** (Please install Docker and docker-compose-plugin manually, and modify .env file)
```shell
# Navigate to /opt and clone the repository
cd /opt
git clone https://github.com/aaPanel/Billion-Mail

# Enter the Billion-Mail directory and copy the file
cd Billion-Mail
cp env_init .env

# Modify the default BILLIONMAIL_HOSTNAME value to your domain, e.g., mail.domain.com
# vi .env

# Initialize the certificate
mkdir ssl
cp -d -n ssl-self-signed/* ssl/

# Start Billion-Mail
docker compose up -d || docker-compose up -d

# Generate bm command
echo "$(pwd)" > /opt/PWD-Billion-Mail.txt
ln -sf $(pwd)/mail_users.sh /usr/bin/bm
chmod +x $(pwd)/mail_users.sh
```


<div align="center">
  <a href="https://www.youtube.com/watch?v=wGHfX1-7S_Y">
    <img src="https://img.youtube.com/vi/wGHfX1-7S_Y/maxresdefault.jpg" alt="" width="80%">
    <br />
    <img src="https://www.iconfinder.com/icons/317714/download/png/16" alt="YouTube" width="16"/>
    <b>Watch on YouTube</b>
  </a>
</div>


## Management script
Management help

`bm help`

View Login default info

`bm default`

Show domain DNS record

`bm show-record`



## WebMail

Billion Mail has integrated **RoundCube**, you can access WebMail via `/roundcube/`.

## Why Billion Mail?

Most email marketing platforms are either **expensive**, **closed-source**, or **lack essential features**. Billion Mail aims to be different:

✅ **Fully Open-Source** – No hidden costs, no vendor lock-in.  
📊 **Advanced Analytics** – Track email delivery, open rates, click-through rates, and more.  
📧 **Unlimited Sending** – No restrictions on the number of emails you can send.  
🎨 **Customizable Templates** – Custom professional marketing templates for reuse.
🔒 **Privacy-First** – Your data stays with you, no third-party tracking.  
🚀 **Self-Hosted** – Run it on your own server for complete control.  

## How You Can Help 🌟

Billion Mail is a **community-driven project**, and we need your support to get started! Here's how you can help:

1. **Star This Repository**: Show your interest by starring this repo.  
2. **Spread the Word**: Share Billion Mail with your network—developers, marketers, and open-source enthusiasts.  
3. **Share Feedback**: Let us know what features you'd like to see in Billion Mail by opening an issue or joining the discussion.  
4. **Contribute**: Once development begins, we'll welcome contributions from the community. Stay tuned for updates!

---

📧 **Billion Mail – The Future of Open-Source Email Marketing.**

## Issues

If you encounter any issues or have feature requests, please [open an issue](https://github.com/your-username/billion-mail/issues). Be sure to include:

- A clear description of the problem or request.
- Steps to reproduce the issue (if applicable).
- Screenshots or error logs (if applicable).

## License

Billion Mail is licensed under the **AGPLv3 License**. This means you can:

✅ Use the software for free.  
✅ Modify and distribute the code.  
✅ Use it privately without restrictions.

See the [LICENSE](LICENSE) file for more details.

---

📬 **Billion Mail – Coming Soon. Star This Repo to Make It Happen Faster!**

<!-- BillionMail official link -->
[docs-link]: https://www.billionmail.com/

<!-- BillionMail Other link-->
[license-link]: https://www.gnu.org/licenses/gpl-3.0.html
[github-release-link]: https://github.com/aaPanel/Billion-Mail/releases/latest
[github-stars-link]: https://github.com/aaPanel/Billion-Mail
[github-issues-link]: https://github.com/aaPanel/Billion-Mail/issues

<!-- Shield link-->
[docs-shield]: https://img.shields.io/badge/documentation-148F76
[github-release-shield]: https://img.shields.io/github/v/release/aaPanel/Billion-Mail
[github-stars-shield]: https://img.shields.io/github/stars/aaPanel/Billion-Mail?color=%231890FF&style=flat-square   
[license-shield]: https://img.shields.io/github/license/aaPanel/Billion-Mail
