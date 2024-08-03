echo "Initiated building parsers for litelang and litequery"

java -jar antlr.jar -Dlanguage=Go -Xexact-output-dir -o build -package LiteLang  -visitor  src/LiteLang.g4 
