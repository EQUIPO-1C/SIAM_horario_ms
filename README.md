# horario_ms

## For developing 
First you have tu create a `.env` file with the next template
    
    DB_USER=<THE DATABASE USER>
    DB_PSWD=<THE DATABASE PASSWORD>

Install  GO and run the next comand in the root directory of the project `go run main.go`

## Using Docker
First build your image in the root directory of the project

    sudo docker build -t schedule_ms_devimg .

Then run your container with the next template

    sudo docker run --network host -e DB_USER=<THE DATABASE USER> -e DB_PSWD=<THE DATABASE PASSWORD> --name schedule_ms schedule_ms_devimg

