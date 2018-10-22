### About

Go library for interacting with the [ElasticEmail API](https://api.elasticemail.com/public/help#start).

### Supported methods

- Email:

    Send your emails and see their statuses

    - [x] GetStatus  
    - [x] Send  
    - [x] Status  
    - [x] View
    
- Template:

    Managing and editing templates of your emails

    - [x] GetList

- List:

    API methods for managing your Lists

    - [x] list
    
- Contact:

    Methods used to manage your Contacts.

    - [x] Add    
        
        
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
