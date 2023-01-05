# pip install pyttsx3
# pip install pypdf2
import pyttsx3
import PyPDF2

#book = input("Enter the name of the PDF file")
book = "test.pdf"
pdf_book = open(book,'rb')
reading_pdf = PyPDF2.PdfFileReader(pdf_book)
pdf_pages = reading_pdf.numPages

read_everything = input("Enter 'Y' if you want to read the entire PDF file, If not enter 'N': ")
if read_everything == 'Y' or read_everything == 'y':
    for i in range(pdf_pages):
        pdf_speaker = pyttsx3.init()
        choose_page = reading_pdf.getPage(i)
        pdf_text = choose_page.extractText()
        pdf_speaker.say(pdf_text)
        pdf_speaker.runAndWait()
else:
    pdf_speaker = pyttsx3.init()
    page = int(input("Enter the name page number: "))
    choose_page = reading_pdf.getPage((page-1))
    pdf_text = choose_page.extractText()
    pdf_speaker.say(pdf_text)
    pdf_speaker.runAndWait()

