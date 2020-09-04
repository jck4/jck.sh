/*
Add 1 letter of the string to an element at a time, on a 10 ms delay. 
*/

function typeEffect(terminalLine, characterPosistion, stringLength, htmlElement) {
  setTimeout(function () {
    if (characterPosistion < stringLength) {
      htmlElement.innerHTML += terminalLine.charAt(characterPosistion++);
      typeEffect(terminalLine, characterPosistion, stringLength, htmlElement); //Recurse
    }
  }, 10);
}


function getShortUrl(url,element) {
fetch('/api/shortenURL',  {
  method: 'POST',
  headers: {
    'Content-Type': undefined
  },

  body: JSON.stringify({
      url
  })

})
  .then((response) => {
    return "ye"
  })
return 'ye'
}

