# Email-Username-Generator
Given a list of names and a list of formats to follow, program will generate a list of emails for each format that follows that given  
i.e. `generator -f formats -n names -o output.txt` creates a file `output.txt` with the formats written to it
## Options
* Use -n to set list of names provided, default is names.csv. __IMPORTANT: format of names must be last name, first name__
* Use -o to set name of output file, default is emails.txt
* Use -f to set name of formats file, default is format.txt. __IMPORTANT: see below for list of formatting options__
* Use -raw to generate a list with only the usernames, without lines and the format of the usernames provided
* Use -d or -duplicates if the list contains duplicates. This will add numbers to the end of the local part of the UN, ie 'jsmith@domain.com' and 'jsmith1@domain.com'  
  
    
## Format List
* (f.) - first initial lowercase
* (F.) - first initial uppercase
* (f) - first name lowercase
* (f~) - first name title case
* (F) - first name capital
* (l.) - last initial lowercase
* (L.) - last initial uppercase
* (l) - last name lowercase
* (l~) - last name title case
* (L) - last name capital  
### Examples (with Smith, John)
* `(f.)(l)@domain.com` -> `jsmith@domain.com`
* `(l~).(f~)@domain.com` -> `Smith.John@domain.com`
* `The_Real_(F)@domain.com` -> `The_Real_JOHN@domain.com`
