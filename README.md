
# Heroku Ip Spoofer
Heroku ip spoofer by tristan
// readme not finished yet
// code not fully optimized yet
<br>

# Usage
Create a new heroku application and commit the code inside the /heroku
to it. This will start the api which you can send requests to.
<br>
Once you're heroku application is up and running, send a request to your
heroku applications url by the following format: 
<br>
String urlToSendRequestTo = "http://httpbin.org"
<br>
HTTP {Method} -> {heroku_url}?url=Base64Encoded(urlToSendRequestTo)

