# Replace date-time declarations with utc-date
if [[ "$OSTYPE" == "darwin"* ]]; then
	sed -i '' 's/"date-time"/"utc-date"/g' tmp/input-schema.json
else
	sed -i 's/"date-time"/"utc-date"/g' tmp/input-schema.json
fi

# Remove uses of the "Search" tag
cat tmp/input-schema.json | jq 'walk( if type == "object" and has("tags") 
	and any(.tags[]; . == "Search") 
	then del(.tags[] | 
	select(. == "Search")) 
	else . end)' > tmp/input-schema.tmp.json

# Remove uses of the "Global rules" tag
cat tmp/input-schema.tmp.json | jq 'walk( if type == "object" and has("tags") 
	and any(.tags[]; . == "Global rules") 
	then del(.tags[] | 
	select(. == "Global rules")) 
	else . end)' > tmp/input-schema.json