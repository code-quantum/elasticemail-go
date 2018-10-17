### About

Implementation of [ElasticEmail API](https://api.elasticemail.com/public/help#start).

### Supported methods

- Email:

    Send your emails and see their statuses

    - [x] GetStatus  
    - [x] Send  
    - [x] Status  
    - [x] View  
### Usage

```go
        ee := elasticemail.NewElasticEmail("api-key")
        
        //send email
	to := make([]string, 1)
	to[0] = "abc@gmail.com"
	sendEmail, err := ee.Send(elasticemail.SendParams{To: to, BodyHtml: "<b>Hello!</b>", Subject: "Hello", From: "noreply@eeee.eee"})
	if err!= nil {
		log.Fatal(err.Error())
	}
	
	// get status:
    	status, err := ee.GetEmailStatus(elasticemail.GetEmailStatusParams{TransactionID: sendEmail.TransactionID})
    	if err != nil {
    		log.Fatal(err.Error())
    	}


```
