# 16ByteChallenge


---

## The Challenge

I've written a dummy login portal, and the challenge is to reverse-engineer it and figure out the correct password to gain access.

- The password is **16 characters long**
- It's **split into 4 equal parts**
- Each part is checked using a different hashing algorithm

If you reverse the logic and enter the correct input, the program will say:

```Access Granted.```

(Or brute-force it if you really want to - you do you) 

---


## Download

You can download the compiled binary from the [Releases](https://github.com/elliott-diy/16ByteChallenge/releases/tag/v1.0.0) section.

---
## Why Go?

I wrote this in Go because I'm a bit of an ass and didn’t realize how much of a pain it would be to reverse.  
Have fun with that.

---

## Resources That Might Help

You probably won’t need all of these, but they might come in handy:

- [Decompiler Explorer](https://dogbolt.org/)
- [NSA’s Ghidra Reverse Engineering Tool](https://github.com/NationalSecurityAgency/ghidra)
- [cyber.elliott.diy](https://cyber.elliott.diy/) — check the "Extract Hashes" option
- [Hashcat](https://hashcat.net/hashcat/) — you don’t need a GPU for these hashes

---

## Building the Binary

To compile:

```bash
go build -o 16ByteChallenge
```
