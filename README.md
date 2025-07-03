# lipi
A Bengali-themed toy programming language — built just for fun, learning, and the joy of creating something cool. Not for production.

# Example Code
## Variable Declaration and Arithmetic
```lipi
ধর সংখ্যা = ৫
ধর নাম = "জারিফ"
ধর যোগফল = সংখ্যা + ১০
ধর ফলাফল = "হ্যালো, " + নাম + "! যোগফল হল: " + যোগফল
বলো(ফলাফল)
```

## Arrays and Maps
```lipi
ধর ফল = ["আপেল", "কলা", "আম"]
ধর ফলাফল = "আমার প্রিয় ফল: " + ফল[০] + ", " + ফল[১] + ", " + ফল[২]
বলো(ফলাফল)

ধর ছাত্র = {
    "নাম": "জারিফ",
    "বয়স": ২৫,
    "কোর্স": "কম্পিউটার সায়েন্স"
}
ধর ছাত্র_নাম = ছাত্র["নাম"]
বলো("ছাত্রের নাম: " + ছাত্র_নাম)
```

## Functions
```lipi
ধর যোগ_করো = ফাঙ্কশন(ক, খ) {
    ফেরাও ক + খ;
};

ধর যোগ_করো = ফাঙ্কশন(ক, খ) {
    ক + খ;
};
```

## Conditionals
```lipi
ধর সংখ্যা = ১
যদি সংখ্যা < ৫ {
    বলো("সংখ্যা ৫ এর চেয়ে ছোট")
} নাহলে {
    বলো("সংখ্যা ৫ এর সমান বা বড়")
}
```