Open TAK Server
===============

 [![Open TAK Server](https://circleci.com/gh/tma5/otaks.svg?style=svg&circle-token=e7d61bafbb0b320b50ae5756b4fa772d19a1989e)](https://app.circleci.com/pipelines/github/tma5/otaks)

FAQ
---

**Does this work yet?**

Nope.

**What's the roadmap?**

There is none so far. Basic functionality is the current goal. 

**Why isn't X implemented?**

I'm not sponsored. I don't have SDK access. I'm hacking on this in spurts of free time. This is purely reverse engineered off public CoT docs, cribbing how [FreeTakServer](https://github.com/Tapawingo/FreeTakServer) is implemented thus far, and slurping down network traffic from ATAK 4.0.0.1.


Quickstart
----------

Start using docker-compose.

```
$ docker-compose up
```


## Refererence

- [Cursor on Target (CoT) Developer Guide, Aug 2005, MITRE](https://apps.dtic.mil/dtic/tr/fulltext/u2/a637348.pdf)
- [Cursor-On-Target Message Router User'S Guide, Nov 2009, MITRE](https://www.mitre.org/sites/default/files/pdf/09_4937.pdf)
- [PHOENIX: Service Oriented Architecture For Information Management The “FAWKES” Cursor-Ontarget Router, Sept 2011, AFRL](https://apps.dtic.mil/dtic/tr/fulltext/u2/a550101.pdf)
- [Yap: A High-Performance Cursor on Target Message Router, Sept 2014, ARL](https://apps.dtic.mil/sti/pdfs/ADA610603.pdf)
