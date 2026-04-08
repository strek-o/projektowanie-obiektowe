package com.example.zadanie03.controllers

import com.example.zadanie03.models.User
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api")
class UserController {

    private val userList = listOf(
        User(1, "test.user01"),
        User(2, "test.user02"),
        User(3, "test.user03")
    )

    @GetMapping("/users")
    fun getUsers(): List<User> {
        return userList
    }

    @GetMapping("/users/{id}")
    fun getUserById(@PathVariable id: Int): ResponseEntity<User> {
        val user = userList.find { it.id == id }
        
        return if (user != null) {
            ResponseEntity.ok(user)
        } else {
            ResponseEntity.notFound().build()
        }
    }
}
