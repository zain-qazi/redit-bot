# redit-bot

# Build Docker image
sudo docker build -t redit-bot-img .

#Check if the image was successfully created
sudo docker image ls

#To inspect the image, run image with bash entrypoint
#This will be helpful in inspecting the container
sudo docker run -it --entrypoint /bin/bash redit-bot-img

#Once you are inside the container, verify if all go env variables look good
go env

#To run the reddit bot 
sudo docker run -it --rm -v ~/go:/app --name redit-bot-running redit-bot redit-bot pull --path="/r/news" --agent-config "/app/bot.agent"

#Command to listen. The path cannot have "/r/"
./redit-bot listen --path="news" --agent-config "/home/zain/go/bot.agent"