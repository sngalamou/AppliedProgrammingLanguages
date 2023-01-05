#pip install pyqt5 
from PyQt5 import QtCore, QtGui, QtWidgets
import sqlite3

tasks = ["Write email", "finish features", "watch tutorial"]

class Ui_MainWindow(object):
    def setupUi(self, MainWindow):
        MainWindow.setObjectName("MainWindow")
        MainWindow.resize(756, 500)
        self.centralwidget = QtWidgets.QWidget(MainWindow)
        self.centralwidget.setObjectName("centralwidget")
        self.tabWidget = QtWidgets.QTabWidget(self.centralwidget)
        self.tabWidget.setGeometry(QtCore.QRect(0, 10, 791, 521))
        self.tabWidget.setMovable(True)
        self.tabWidget.setTabBarAutoHide(False)
        self.tabWidget.setObjectName("tabWidget")
        self.planner = QtWidgets.QWidget()
        self.planner.setObjectName("planner")
        self.calendarWidget = QtWidgets.QCalendarWidget(self.planner)
        self.calendarWidget.setGeometry(QtCore.QRect(0, 30, 411, 341))
        self.calendarWidget.setStyleSheet("position: left")
        self.calendarWidget.setObjectName("calendarWidget")
        self.calendarWidget.selectionChanged.connect(self.calendarDateChanged)
        self.listWidget = QtWidgets.QListWidget(self.planner)
        self.listWidget.setGeometry(QtCore.QRect(450, 90, 281, 251))
        self.listWidget.setObjectName("listWidget")
        self.updateTaskList()
        self.pushButton = QtWidgets.QPushButton(self.planner)
        self.pushButton.setGeometry(QtCore.QRect(480, 350, 241, 23))
        self.pushButton.setStyleSheet("border-radius:10px;\n"
"background-color: red;\n"
"color: white;\n"
"font: 11pt;")
        self.pushButton.setObjectName("pushButton")
        self.pushButton_2 = QtWidgets.QPushButton(self.planner)
        self.pushButton_2.setGeometry(QtCore.QRect(650, 60, 75, 23))
        self.pushButton_2.setStyleSheet("border-radius:10px;\n"
"background-color: blue;\n"
"color: white;\n"
"font: 12pt;")
        self.pushButton_2.setObjectName("pushButton_2")
        self.tabWidget.addTab(self.planner, "")
        self.pomodoro = QtWidgets.QWidget()
        self.pomodoro.setObjectName("pomodoro")
        self.pushButton_6 = QtWidgets.QPushButton(self.pomodoro)
        self.pushButton_6.setGeometry(QtCore.QRect(490, 340, 121, 31))
        self.pushButton_6.setStyleSheet("border-radius:10px;\n"
"background-color: blue;\n"
"color: white;\n"
"font: 11pt;\n"
"")
        self.pushButton_6.setObjectName("pushButton_6")
        self.pushButton_7 = QtWidgets.QPushButton(self.pomodoro)
        self.pushButton_7.setGeometry(QtCore.QRect(350, 330, 121, 21))
        self.pushButton_7.setStyleSheet("border-radius:10px;\n"
"background-color: red;\n"
"color: white;\n"
"font: 11pt;\n"
"")
        self.pushButton_7.setObjectName("pushButton_7")
        self.pushButton_8 = QtWidgets.QPushButton(self.pomodoro)
        self.pushButton_8.setGeometry(QtCore.QRect(210, 340, 121, 31))
        self.pushButton_8.setStyleSheet("border-radius:10px;\n"
"background-color: green;\n"
"color: white;\n"
"font: 11pt;\n"
"")
        self.pushButton_8.setObjectName("pushButton_8")
        self.lcdNumber = QtWidgets.QLCDNumber(self.pomodoro)
        self.lcdNumber.setGeometry(QtCore.QRect(240, 90, 281, 121))
        self.lcdNumber.setObjectName("lcdNumber")
        self.pushButton_9 = QtWidgets.QPushButton(self.pomodoro)
        self.pushButton_9.setGeometry(QtCore.QRect(350, 360, 121, 21))
        self.pushButton_9.setStyleSheet("border-radius:10px;\n"
"background-color: red;\n"
"color: white;\n"
"font: 11pt;\n"
"")
        self.pushButton_9.setObjectName("pushButton_9")
        self.tabWidget.addTab(self.pomodoro, "")
        self.tts = QtWidgets.QWidget()
        self.tts.setObjectName("tts")
        self.pushButton_3 = QtWidgets.QPushButton(self.tts)
        self.pushButton_3.setGeometry(QtCore.QRect(210, 340, 121, 31))
        self.pushButton_3.setStyleSheet("border-radius:10px;\n"
"background-color: green;\n"
"color: white;\n"
"font: 11pt;\n"
"")
        self.pushButton_3.setObjectName("pushButton_3")
        self.label = QtWidgets.QLabel(self.tts)
        self.label.setGeometry(QtCore.QRect(10, 30, 121, 81))
        self.label.setObjectName("label")
        self.pushButton_4 = QtWidgets.QPushButton(self.tts)
        self.pushButton_4.setGeometry(QtCore.QRect(490, 340, 121, 31))
        self.pushButton_4.setStyleSheet("border-radius:10px;\n"
"background-color: blue;\n"
"color: white;\n"
"font: 11pt;\n"
"")
        self.pushButton_4.setObjectName("pushButton_4")
        self.pushButton_5 = QtWidgets.QPushButton(self.tts)
        self.pushButton_5.setGeometry(QtCore.QRect(350, 340, 121, 31))
        self.pushButton_5.setStyleSheet("border-radius:10px;\n"
"background-color: red;\n"
"color: white;\n"
"font: 11pt;\n"
"")
        self.pushButton_5.setObjectName("pushButton_5")
        self.tabWidget.addTab(self.tts, "")
        MainWindow.setCentralWidget(self.centralwidget)
        self.menubar = QtWidgets.QMenuBar(MainWindow)
        self.menubar.setGeometry(QtCore.QRect(0, 0, 756, 21))
        self.menubar.setObjectName("menubar")
        self.menuFile = QtWidgets.QMenu(self.menubar)
        self.menuFile.setObjectName("menuFile")
        self.menuPomodoro = QtWidgets.QMenu(self.menubar)
        self.menuPomodoro.setObjectName("menuPomodoro")
        self.menuHelp = QtWidgets.QMenu(self.menubar)
        self.menuHelp.setObjectName("menuHelp")
        MainWindow.setMenuBar(self.menubar)
        self.statusbar = QtWidgets.QStatusBar(MainWindow)
        self.statusbar.setObjectName("statusbar")
        MainWindow.setStatusBar(self.statusbar)
        self.toolBar = QtWidgets.QToolBar(MainWindow)
        self.toolBar.setObjectName("toolBar")
        MainWindow.addToolBar(QtCore.Qt.TopToolBarArea, self.toolBar)
        self.actionTimer_Start = QtWidgets.QAction(MainWindow)
        icon = QtGui.QIcon()
        icon.addPixmap(QtGui.QPixmap("icons/alarm-clock--plus.png"), QtGui.QIcon.Normal, QtGui.QIcon.Off)
        self.actionTimer_Start.setIcon(icon)
        self.actionTimer_Start.setObjectName("actionTimer_Start")
        self.actionTimer_End = QtWidgets.QAction(MainWindow)
        icon1 = QtGui.QIcon()
        icon1.addPixmap(QtGui.QPixmap("icons/alarm-clock--minus.png"), QtGui.QIcon.Normal, QtGui.QIcon.Off)
        self.actionTimer_End.setIcon(icon1)
        self.actionTimer_End.setObjectName("actionTimer_End")
        self.actionNew = QtWidgets.QAction(MainWindow)
        self.actionNew.setObjectName("actionNew")
        self.menuFile.addAction(self.actionNew)
        self.menuPomodoro.addAction(self.actionTimer_Start)
        self.menuPomodoro.addAction(self.actionTimer_End)
        self.menubar.addAction(self.menuFile.menuAction())
        self.menubar.addAction(self.menuPomodoro.menuAction())
        self.menubar.addAction(self.menuHelp.menuAction())
        self.toolBar.addAction(self.actionTimer_Start)
        self.toolBar.addAction(self.actionTimer_End)

        self.retranslateUi(MainWindow)
        self.tabWidget.setCurrentIndex(1)
        QtCore.QMetaObject.connectSlotsByName(MainWindow)

    def updateTaskList(self):
            for task in tasks:
                    #add item to the list widget
                    item = QtWidgets.QListWidgetItem(task)
                    item.setFlags(item.flags() | QtCore.Qt.ItemIsUserCheckable)
                    item.setCheckState(QtCore.Qt.Unchecked)
                    self.listWidget.addItem(item)
    def calendarDateChanged(self):
        print("CalendarDateChanged")
        dateSelected = self.calendarWidget.selectedDate().toPyDate().strftime("%Y-%m-%d")
        print("Date Selected: ", dateSelected)
    
    def retranslateUi(self, MainWindow):
        _translate = QtCore.QCoreApplication.translate
        MainWindow.setWindowTitle(_translate("MainWindow", "The Study Tool"))
        self.pushButton.setText(_translate("MainWindow", "Save Changes"))
        self.pushButton_2.setText(_translate("MainWindow", "Add New"))
        self.tabWidget.setTabText(self.tabWidget.indexOf(self.planner), _translate("MainWindow", "Planner"))
        self.pushButton_6.setText(_translate("MainWindow", "Stop Timer"))
        self.pushButton_7.setText(_translate("MainWindow", "30 Minute Timer"))
        self.pushButton_8.setText(_translate("MainWindow", "Start Timer"))
        self.pushButton_9.setText(_translate("MainWindow", "45 Minute Timer"))
        self.tabWidget.setTabText(self.tabWidget.indexOf(self.pomodoro), _translate("MainWindow", "Pomodoro"))
        self.pushButton_3.setText(_translate("MainWindow", "Play"))
        self.label.setText(_translate("MainWindow", "Now Reading: "))
        self.pushButton_4.setText(_translate("MainWindow", "Pause"))
        self.pushButton_5.setText(_translate("MainWindow", "Stop"))
        self.tabWidget.setTabText(self.tabWidget.indexOf(self.tts), _translate("MainWindow", "Text-To-Speech"))
        self.menuFile.setTitle(_translate("MainWindow", "File"))
        self.menuPomodoro.setTitle(_translate("MainWindow", "Pomodoro"))
        self.menuHelp.setTitle(_translate("MainWindow", "Help"))
        self.toolBar.setWindowTitle(_translate("MainWindow", "toolBar"))
        self.actionTimer_Start.setText(_translate("MainWindow", "Timer Start"))
        self.actionTimer_End.setText(_translate("MainWindow", "Timer End"))
        self.actionNew.setText(_translate("MainWindow", "Open"))


if __name__ == "__main__":
    import sys
    app = QtWidgets.QApplication(sys.argv)
    MainWindow = QtWidgets.QMainWindow()
    ui = Ui_MainWindow()
    ui.setupUi(MainWindow)
    MainWindow.show()
    sys.exit(app.exec_())
