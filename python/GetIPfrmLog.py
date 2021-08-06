#Progam to Get Unique IPaddress from the String of Logs
import re
test_string = """110.110.110.10 testing one 120.120.330.120 heelo"""  
list = test_string.split()
print(list)
pattern = re.compile(r'(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})')
newlist = filter(pattern.match, list) 
for i in set(newlist):
#for i in newlist:
    print(i)
