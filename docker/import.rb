require 'json'

import_file='import.json'
if not File.exist?(import_file)
	exit
end

File.open(import_file, 'r') {|f|
	JSON.load(f.read()).each{ |endpoint|
		tmpfile=import_file + ".tmp"
		File.open(tmpfile, 'w') {|write_data|
			write_data.write(JSON.generate(endpoint)) 
		}
		`curl -H "Content-Type: application/json" -d @#{tmpfile} -v -g "http://localhost:8080/endpoints"`
		File.delete(tmpfile)
	}
}

