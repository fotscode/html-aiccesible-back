#!/bin/sh
docker exec -it ollama ollama serve&
docker exec -it ollama ollama create html-analyst-llama3.1 -f /models/Modelfile_llama3.1 
