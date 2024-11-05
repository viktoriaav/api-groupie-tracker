# API groupie tracker

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information. 

## How to run

- Clone the repo 
```
git clone https://01.kood.tech/git/anderterehov/groupie-tracker-visualizations.git
```

- run the program
```
go run . 
```
- open this URL in your browser http://localhost:8080/

- CTRL+C to quit the program

## Implementation details

Reading info from artists API into a slice of structs. Displaying the image and artists names as cards on the homepage. Cards are clickable and are calling popup with additional information about chosen artist.
Close button added as well to to close the artists pages. 

## Authors
- [Ander](https://01.kood.tech/git/anderterehov)
- [Viktoriia](https://01.kood.tech/git/vavstanc)

