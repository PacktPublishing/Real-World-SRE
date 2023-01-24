


# Real-World SRE
This is the code repository for [Real-World SRE](https://www.packtpub.com/web-development/real-world-sre?utm_source=github&utm_medium=repository&utm_campaign=9781788628884), published by [Packt](https://www.packtpub.com/?utm_source=github). It contains all the supporting project files necessary to work through the book from start to finish.
## About the Book
Real-World SRE is the go-to survival guide for the software developer in the middle of catastrophic website failure. Site Reliability Engineering (SRE) has emerged on the frontline as businesses strive to maximize uptime. This book is a step-by-step framework to follow when your website is down and the countdown is on to fix it.

Nat Welch has battle-hardened experience in reliability engineering at some of the biggest outage-sensitive companies on the internet. Arm yourself with his tried-and-tested methods for monitoring modern web services, setting up alerts, and evaluating your incident response.

Real-World SRE goes beyond just reacting to disaster—uncover the tools and strategies needed to safely test and release software, plan for long-term growth, and foresee future bottlenecks. Real-World SRE gives you the capability to set up your own robust plan of action to see you through a company-wide website crisis.

The final chapter of Real-World SRE is dedicated to acing SRE interviews, either in getting a first job or a valued promotion.
## Instructions and Navigation
All of the code is organized into folders. Each folder starts with a number followed by the application name. For example, Chapter02.



The code will look like the following:
```
not_found do
  $statsd.time "request.time" do
    $statsd.increment "request.count"
    $statsd.increment "request.error"
    "This is not found."
  end
end
```

All Go code in this book is assuming Go version 1.10. You can get the latest version from https://golang.org/dl/.

All Ruby code in this book assumes Ruby 2.5, which can be acquired from https://www.ruby-lang.org/en/downloads/releases/.

## Related Products
* [Hands-On Security in DevOps](https://www.packtpub.com/networking-and-servers/hands-security-devops?utm_source=github&utm_medium=repository&utm_campaign=9781788995504)

* [Practical DevOps - Second Edition](https://www.packtpub.com/virtualization-and-cloud/practical-devops-second-edition?utm_source=github&utm_medium=repository&utm_campaign=9781788392570)

* [DevOps with Kubernetes](https://www.packtpub.com/virtualization-and-cloud/devops-kubernetes?utm_source=github&utm_medium=repository&utm_campaign=9781788396646)

### Download a free PDF

 <i>If you have already purchased a print or Kindle version of this book, you can get a DRM-free PDF version at no cost.<br>Simply click on the link to claim your free PDF.</i>
<p align="center"> <a href="https://packt.link/free-ebook/9781788628884">https://packt.link/free-ebook/9781788628884 </a> </p>