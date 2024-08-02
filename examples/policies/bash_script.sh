# Rename all *.txt to *.text
for file in *.go; do 
    mv -- "$file" "${file%.go}.md"
done
