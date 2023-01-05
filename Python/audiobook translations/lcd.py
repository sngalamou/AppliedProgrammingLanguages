# pip install pyqt5

from PyQt5.QtWidgets import QMainWindow, QApplication
from PyQt5 import uic
import sys

class UI(QMainWindow):
    def __init__(self):
        super(UI, self).__init__()

        #load the ui file
        uic.loadUi("lcd.ui",self)

        #define our widgets

        #show the app
        self.show()

#Initialize the app
app = QtWidgets.QApplication(sys.argv)
UIWindow = UI()
app.exec_()