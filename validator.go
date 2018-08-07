package main


import "regexp"
import "strconv"
import "net/http"
import "io"
import "html/template"


var vfrm *template.Template

	func init(){//defines the files that are ready to go for the template
	vfrm = template.Must(template.ParseGlob("vatForm.gohtml"))	
	}

func main() {
	http.HandleFunc("/", vatForm)
	http.HandleFunc("/process",processor)
	http.ListenAndServe(":8080",nil) //Listens and serves on port 8080
}

func vatForm(w http.ResponseWriter, r*http.Request){//Response Writer prints output to html template
	
	vfrm.ExecuteTemplate(w,"vatForm.gohtml",nil)
}

func processor(w http.ResponseWriter, r*http.Request){//All countries' prefix and VAT numbers are processed and matched with input
	if r.Method != "POST"{ 
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
		}
		prefix := r.FormValue("pre")
		number := r.FormValue("num")
		io.WriteString(w, prefix)
		io.WriteString(w, number)
		io.WriteString(w, " is ")
		switch prefix {
		case "AT":
			match, _ := regexp.MatchString("(AT)?U[0-9]{8}$",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
			break	
		case "BE":
			match, _ := regexp.MatchString("BE([0-1]{1}[0-9]{9}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "BG":
			match, _ := regexp.MatchString("BG([0-9]{9,10}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "HR":
			match, _ := regexp.MatchString("HR([0-9]{11}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "CY":
			match, _ := regexp.MatchString("CY([A-Z0-9]{9}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "CZ":
			match, _ := regexp.MatchString("CZ([0-9]{8,10}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "DK":
			match, _ := regexp.MatchString("DK([0-9]{8}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "EE":
			match, _ := regexp.MatchString("EE([0-9]{9}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "FI":
			match, _ := regexp.MatchString("FI([0-9]{8}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "DE":
			match, _ := regexp.MatchString("DE([0-9]{9}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "HU":
			match, _ := regexp.MatchString("HU([0-9]{8}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "IT":
			match, _ := regexp.MatchString("IT([0-9]{11}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "LV":
			match, _ := regexp.MatchString("LV([0-9]{11}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "LT":
			match, _ := regexp.MatchString("LT([0-9]{9,12}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "LU":
			match, _ := regexp.MatchString("LU([0-9]{8}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "MT":
			match, _ := regexp.MatchString("MT([0-9]{8}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "NL":
			match, _ := regexp.MatchString("NL([0-9]{9}B[0-9]{2}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "PL":
			match, _ := regexp.MatchString("PL([0-9]{10}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "PT":
			match, _ := regexp.MatchString("PT([0-9]{9}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "RO":
			match, _ := regexp.MatchString("RO([0-9]{2,10}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "SI":
			match, _ := regexp.MatchString("SI([0-9]{8}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "ES":
			match, _ := regexp.MatchString("ES([A-Z]{1}[0-9]{8}$)",prefix+number)
			match1, _ := regexp.MatchString("ES([A-Z]{1}[0-9]{7}[A-Z]{1}$)",prefix+number)
			if match == true || (match1 == true){
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "SE":
			match, _ := regexp.MatchString("SE([0-9]{12}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "EL":
			match, _ := regexp.MatchString("EL([0-9]{9}$)",prefix+number)
			
			if match == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "GR":
			match, _ := regexp.MatchString("GR([0-9]{9}$)",prefix+number)
			
			if match  == true{
				io.WriteString(w,"Valid")
			}else{
				io.WriteString(w,"Invalid")
			}
		case "IE":
			match, _ := regexp.MatchString("IE([0-9]{7}[A-Z]{2}$)",prefix+number)
			match1, _ := regexp.MatchString("IE([0-9]{7}[A-Z]{1}$)",prefix+number)
		        if (match == true) || (match1 == true){
				io.WriteString(w,"Valid")	
			}else{
				io.WriteString(w,"Invalid")
			}
		case "SK":
			match, _ := regexp.MatchString("SK([0-9]{10}$)",prefix+number)
			if match == true{
				resultnumber, _ := strconv.ParseInt(number,10,32)
				if(resultnumber%11==0){
				io.WriteString(w,"Valid")
				}else{
				io.WriteString(w,"Invalid")	
				}
			   }else{
				io.WriteString(w,"Invalid")
			}
		case "GB":
			match, _ := regexp.MatchString("GB([0-9]{9}$)",prefix+number)
			match1, _ := regexp.MatchString("GB([0-9]{12}$)",prefix+number)
			match2, _ := regexp.MatchString("GB((GD)0*([0-9]|[1-8][0-9]|9[0-9]|[1-3][0-9]{2}|4[0-8][0-9]|49[0-9])$)",prefix+number)
			match3, _ := regexp.MatchString("GB((HA)0*([5-8][0-9]{2}|9[0-8][0-9]|99[0-9])$)",prefix+number)
		        if (match == true) || (match1 == true) || (match2 == true) || (match3 == true) {
				io.WriteString(w,"Valid")	
			}else{
				io.WriteString(w,"Invalid")
			}
		case "FR":
			match, _ := regexp.MatchString("FR([0-9]{11}$)",prefix+number)
			match1, _ := regexp.MatchString("FR([A-Z]{2}[0-9]{9}$)",prefix+number)
			if match == true{
				x1 := number[0:2]
				x2 := number[2:11]
				x11, _ := strconv.ParseInt(x1,10,32)
				x22, _ := strconv.ParseInt(x2,10,64)
				formula := ((12+(3*(x22%97)))%97)
				if (formula==x11){
				io.WriteString(w,"Valid")
				}else{
				io.WriteString(w,"Invalid")	
				}
		
			}else{ 
				if match1 == true{
					io.WriteString(w,"Valid")
				 }else{
					io.WriteString(w,"Invalid")
				}
			}
		default:
			io.WriteString(w,"Invalid")	
		

	}
		
}

