import data.MessageData
import java.io.IOException
import java.net.SocketTimeoutException
import java.net.URI
import java.net.URLEncoder
import java.net.http.HttpClient
import java.net.http.HttpRequest
import java.net.http.HttpResponse

object MainViewModel {
    private val client = HttpClient.newBuilder().build()
    private const val URL = "http://localhost:3333"

    // TODO: if the server is not found, suggest to change server ip
    fun sendMessage(messageData: MessageData) {
        val data = mapOf(
            "message" to """
            Name: ${messageData.name}
            Issue: ${messageData.message}
        """.trimIndent()
        )

        val request = HttpRequest.newBuilder().uri(URI.create("$URL/sendSMS")).POST(formData(data))
            .header("Content-Type", "application/x-www-form-urlencoded").build()
        try {
            val response = client.send(request, HttpResponse.BodyHandlers.ofString())
            when(response.statusCode()) {
                500 -> {
                    // TODO: Show a dialog where sending failed
                    println("Sending failed with error: ${response.body()}")
                }
                200 -> {
                    println("Sent SMS: ${response.statusCode()}")
                }
            }
        } catch (e: SocketTimeoutException) {
            println(e.message)
        }
    }

    private fun String.utf8(): String = URLEncoder.encode(this, "UTF-8")

    private fun formData(data: Map<String, String>): HttpRequest.BodyPublisher? {
        val res = data.map { (k, v) -> "${(k.utf8())}=${v.utf8()}" }
            .joinToString("&")

        return HttpRequest.BodyPublishers.ofString(res)
    }
}