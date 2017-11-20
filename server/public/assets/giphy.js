const API_KEY = 'LdS8wjf7ZiQWV9HHeWMbXwQt31tiJ0xY'
const giphyAPI = `https://api.giphy.com/v1/gifs/search?q=404&api_key=${API_KEY}`
const app = document.querySelector('#app')
const image = document.createElement('img')

app.append(image)

appendGif()

function appendGif() {
  fetch(giphyAPI)
    .then(response => response.json())
    .then(gifs => {
      image.src = gifs.data[0].images.original.url
      loopGifs(gifs.data, 1)
    })
}

function loopGifs(gifs, index) {
  const i = gifs.length <= index ? 0 : index + 1
  image.src = gifs[index].images.original.url
  setTimeout(loopGifs.bind(null, gifs, i), 5000)
}
