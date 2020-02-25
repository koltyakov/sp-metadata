# Namespace: Microsoft.Office.Server.Search.REST

## Function Import: suggest

- Entity type: Microsoft.Office.Server.Search.REST.SearchService
- Is composable: false
- Is bindable: false

### Parameters

Parameter | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------
querytext (Edm.String) | ✅ | ✅ | ✅ | ✅
iNumberOfQuerySuggestions (Edm.Int32) | ✅ | ✅ | ✅ | ✅
iNumberOfResultSuggestions (Edm.Int32) | ✅ | ✅ | ✅ | ✅
iNumberOfPopularResultSuggestions (Edm.Int32) | ✅ | ❌ | ❌ | ❌
fPreQuerySuggestions (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
fHitHighlighting (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
fCapitalizeFirstLetters (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
culture (Edm.Int32) | ✅ | ✅ | ✅ | ✅
enableStemming (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
showPeopleNameSuggestions (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
enableQueryRules (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
fPrefixMatchAllTerms (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
sourceId (Edm.String) | ✅ | ✅ | ❌ | ✅
clientType (Edm.String) | ✅ | ✅ | ❌ | ❌
useOLSQuery (Edm.Int32) | ✅ | ✅ | ❌ | ❌
OLSQuerySession (Edm.String) | ✅ | ✅ | ❌ | ❌
zeroTermSuggestions (Edm.Boolean) | ✅ | ✅ | ❌ | ❌
