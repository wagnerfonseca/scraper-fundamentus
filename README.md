# Scraper Fundamentus
This code is used to extract data from [fundamentus.com.br](http://www.fundamentus.com.br) website.
Get all financial and fundamentalists data from companies listed on the [Bovespa](http://www.b3.com.br/pt_br/).
​
## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
​
### Prerequisites
Stuff that you need to install the software and how to run them
​
```
Go 1.12 or above
```
​
### Running
Run the command below on your local machine (works better on Linux/macOS terminal emulation).
​
```
make run
```
​
When the script finishes running, a csv file naming `fundamentus.csv` will be ready with all fundamentus website data.

## Built With

* [colly](http://go-colly.org/) - The scraping framework used

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Authors

* **Wagner Fonseca** - *Initial work* - [wagnerfonseca](https://github.com/wagnerfonseca)
* **Hurick Krügner** - *Review* - [hurick](https://github.com/hurick)
