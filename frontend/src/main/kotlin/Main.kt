import androidx.compose.desktop.ui.tooling.preview.Preview
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.compose.ui.window.application
import data.MessageData
import org.jetbrains.jewel.foundation.theme.JewelTheme
import org.jetbrains.jewel.intui.standalone.theme.IntUiTheme
import org.jetbrains.jewel.intui.standalone.theme.darkThemeDefinition
import org.jetbrains.jewel.intui.standalone.theme.default
import org.jetbrains.jewel.intui.window.decoratedWindow
import org.jetbrains.jewel.intui.window.styling.dark
import org.jetbrains.jewel.ui.ComponentStyling
import org.jetbrains.jewel.ui.component.OutlinedButton
import org.jetbrains.jewel.ui.component.Text
import org.jetbrains.jewel.ui.component.TextArea
import org.jetbrains.jewel.ui.component.TextField
import org.jetbrains.jewel.window.DecoratedWindow
import org.jetbrains.jewel.window.DecoratedWindowScope
import org.jetbrains.jewel.window.TitleBar
import org.jetbrains.jewel.window.newFullscreenControls
import org.jetbrains.jewel.window.styling.TitleBarColors
import org.jetbrains.jewel.window.styling.TitleBarStyle

const val APP_NAME = "TechAlert"
val accent = Color(248, 159, 34)
val onAccent = Color(160, 31, 9)

@Composable
fun DecoratedWindowScope.AppBar() {
    TitleBar(
        Modifier.newFullscreenControls(),
        style = TitleBarStyle.dark(
            colors = TitleBarColors.dark(
                backgroundColor = accent,
                borderColor = accent,
                inactiveBackground = accent
            )
        )
    ) {
        Text(
            "Cagayan State-University",
//            fontSize = 30.sp,
            color = onAccent,
//            fontWeight = FontWeight.SemiBold,
//            modifier = Modifier.padding(20.dp)
        )
    }
}

@Composable
fun EditText(label: String, value: String, onValueChange: (String) -> Unit, modifier: Modifier = Modifier) {
    Row(verticalAlignment = Alignment.CenterVertically) {
        Text(label, fontSize = 15.sp)
        TextField(
            modifier = modifier,
            value = value,
            onValueChange = { onValueChange(it) }
        )
    }
}

@Composable
@Preview
fun App(exitApplication: () -> Unit) {
    IntUiTheme(
        theme = JewelTheme.darkThemeDefinition(),
        styling = ComponentStyling.default().decoratedWindow(titleBarStyle = TitleBarStyle.dark()),
    ) {
        DecoratedWindow(
            title = APP_NAME,
            onCloseRequest = { exitApplication() },
            content = {
                AppBar()

                // TODO: add a default lab room
                var laboratoryRoom by remember { mutableStateOf("") }
                var pcNumber by remember { mutableStateOf("") }
                var name by remember { mutableStateOf("") }
                var yearSection by remember { mutableStateOf("") }
                var professor by remember { mutableStateOf("") }
                var concern by remember { mutableStateOf("") }

                val mainUiHeightRatio = .8f

                Column(Modifier.fillMaxSize().background(color = JewelTheme.globalColors.panelBackground)) {
//                    Box(Modifier.background(accent).fillMaxWidth(), contentAlignment = Alignment.Center) {
//                    }
                    Row(
                        Modifier.fillMaxHeight(),
                        verticalAlignment = Alignment.CenterVertically
                    ) {
                        Column(
                            Modifier.weight(1f).fillMaxHeight(mainUiHeightRatio).padding(20.dp),
                            verticalArrangement = Arrangement.SpaceBetween
                        ) {
                            val textSpace = 20.dp

                            Column(verticalArrangement = Arrangement.spacedBy(textSpace)) {
                                Spacer(Modifier.height(20.dp))

                                EditText(
                                    "Laboratory Room: ",
                                    laboratoryRoom,
                                    onValueChange = { laboratoryRoom = it },
                                    Modifier.fillMaxWidth()
                                )
                                EditText(
                                    "PC Number: ",
                                    pcNumber,
                                    onValueChange = { pcNumber = it },
                                    Modifier.fillMaxWidth()
                                )
                            }

                            Column(verticalArrangement = Arrangement.spacedBy(textSpace)) {
                                EditText("Name: ", name, onValueChange = { name = it }, Modifier.fillMaxWidth())
                                EditText(
                                    "Year/Section: ",
                                    yearSection,
                                    onValueChange = { yearSection = it },
                                    Modifier.fillMaxWidth()
                                )
                                EditText(
                                    "Professor",
                                    professor,
                                    onValueChange = { professor = it },
                                    Modifier.fillMaxWidth()
                                )

                                Spacer(Modifier.height(25.dp))
                            }
                        }
                        Column(
                            Modifier.weight(1f).fillMaxHeight(mainUiHeightRatio).padding(20.dp),
                            verticalArrangement = Arrangement.spacedBy(20.dp)
                        ) {
                            Text("Put your concern here", fontSize = 15.sp, fontStyle = FontStyle.Italic)
                            TextArea(
                                value = concern,
                                onValueChange = { concern = it },
                                modifier = Modifier.weight(1f).fillMaxWidth()
                            )
                            Box(Modifier.fillMaxWidth(), contentAlignment = Alignment.Center) {
                                OutlinedButton(onClick = {
                                    MainViewModel.sendMessage(MessageData(name, concern))
                                    println("hello")
                                }, content = { Text("Send") })
                            }
                        }
                    }
                }
            }
        )
    }
}

fun main() = application {
    App { exitApplication() }
}
