package datasource

import (
	"net/http"
	"net/url"
	"fmt"
	"log"
	"io/ioutil"
)

func GetDataNorge(host string,source string ,longitude string ,latitude string , radius string , date string , school string)([]byte,error){
	params := fmt.Sprintf("source=%s&longitude=%s&latitude=%s&radius=%s&dato=%s&skole=%s",source,longitude,latitude,radius,date,url.QueryEscape(school))
	url := host+"/DataNorge?"+params
	//url := "http://lego.fiicha.net:8080/DataNorge?"+params
	log.Println(url)
	resp,err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		return nil,err
	}else{
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		contents_str := string(contents)
		if err != nil {
		    log.Printf("%s", err)
		    return nil,err
		}else {
		    log.Printf("%s\n",contents_str )
		    return contents , nil
		}
	}

}
