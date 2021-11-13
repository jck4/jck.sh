function generatePass(){
    var keyAlpha  ="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
    keyInt= "123456789",
    keySpecial= '!\"#$%&\'()*+,-./:;<=>?@[\\]^_`{|}~';
    keys =[keyInt,keySpecial,keyAlpha]
    
    keys[Math.floor(Math.random() * 3)]

    var passLen = document.getElementById("passLength").value;
    var password = ""
    //If the password isn't an integer
    if (passLen > 100 || passLen <=7 ){return 1}
    for (let i = 0; i < passLen; i++) { 
        keyType = keys[Math.floor(Math.random() * 3)]
        password += keyType[Math.floor(Math.random() * keyType.length)]
    }
    var passwordBox = document.getElementById("passwordBox");
    passwordBox.innerText = password;
}