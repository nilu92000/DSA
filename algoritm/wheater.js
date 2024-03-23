console.log('hey');
const API_KEY="d1845658f92b31c64bd94f06f7188c9c";
async function getCustomWeatherDetail(){
try{
    let city="goa";
    const response=await fetch('https://api.openweathermap.org/data/2.5/weateher?=${city}&appid=${API_KEY}&units=metric');
    const data=await response.json();
    console.log("weather data:->",data);
    let para=document.createElement('p');
    newpara.textContent='${data?.main?.temp.toFixed(2)}°C';
    document.body.appendChild(newpara);
    randerWeartherInfo(data);
}
catch(err){

}
//https://api.openweathermap.org/data/2.5/weateher?=$%7Bcity%7D&appid=$%7BAPI_KEY%7D&units=metric
}
async function getCustomWeatherDetails(){
    let latitute=15.6333;
    let longitute=18.6333;
    let result=await fetch('https://api.openweathermap.org/data/2.5/weateher?=$%7Bcity%7D&appid=$%7BAPI_KEY%7D&units=metric');
    let data=await result.json;
    console.log(data);
}
catch(err){
    console.log("error found",err);
}

