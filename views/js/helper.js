function typeEffect (data,index,initial,element) {
    setTimeout(function () {
      //Add 1 letter of the string to an element at a time
      if (index < initial) {         
        element.innerHTML += data.charAt(index++);
        typeEffect(data, index, initial,element); //Recurse
      }
    }, 10);
  }