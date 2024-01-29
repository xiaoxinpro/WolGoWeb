Dim objFSO, objFile, strContent, objRegEx, args, adodbStream

Set objFSO = CreateObject("Scripting.FileSystemObject")
Set args = WScript.Arguments
Set adodbStream = CreateObject("ADODB.Stream")

adodbStream.Charset = "utf-8"
adodbStream.Open
adodbStream.LoadFromFile args(0)
strContent = adodbStream.ReadText
adodbStream.Close

Set objRegEx = New RegExp
objRegEx.Global = True
objRegEx.IgnoreCase = True
objRegEx.Pattern = "VERSION\s+=\s+\""[\d\.]+\"""
strContent = objRegEx.Replace(strContent, "VERSION = """ & args(1) & """")

adodbStream.Charset = "utf-8"
adodbStream.Open
adodbStream.WriteText strContent
adodbStream.SaveToFile args(0), 2
adodbStream.Close