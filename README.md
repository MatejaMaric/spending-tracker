**I stopped working on this project because now I'm using GnuCash. I found out
that my banks supports exporting transactions into CSV, which GnuCash is able
to import. Although GnuCash is kind of clunky it does everything I need and
more. On the positive note, I learned quite a bit from this project and I have
a strange feeling that this code will prove useful in the future.**

# About this project:

I'm making this app to help me easily label and plot my transactions.
It's made using Go, SQLite and Fyne (a cross-platform GUI framework for Go).

My bank doesn't have open API endpoints for me to use, so I need to scrape HTML tables from their web interface...
I basically go to my bank's website, login into my account, open dev tools and copy the outer HTML of the transaction table.
Then I paste the HTML into my application that parses and stores the transactions for me.

Since you almost surly don't use the same bank as I do, this app is useless for you.
Unless, that is, you have the same problem as I do but with a different bank.
Then you can pretty easily modify the parser to suit your needs.

## How to run this project:

Since this project uses Fyne, you will need to install libraries it requires for your platform. 
You can find them in [Fyne Documentation](https://developer.fyne.io/started/#prerequisites).

```
go build -v
./spending-tracker
```

## How to try this application:

There is an `example_table` file from which you can copy HTML to insert into this app...
