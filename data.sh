COUNTER="1"

while true 
do
    temp=$(shuf -i 18-23 -n 1);
    hum=$(shuf -i 70-100 -n 1);
    co2=$(shuf -i 40-70 -n 1);
     curl localhost:8080/data -S -H 'Content-Type: application/json' -X POST -d '{"timestamp": "45","id":"MOD1","mac": "12AB12CD23","lat": 12.34,"long": 12.34,"params": {"temperature":'$temp',"humidity" :'$hum',"co2" :'$co2'}}'


    sleep 2.5


done



