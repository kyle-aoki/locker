package com.secretsLocker.locker;

import com.secretsLocker.locker.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.Banner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class LockerApplication {

	@Autowired
	static UserRepository userRepository;

	public static void main(String[] args) {
		SpringApplication.run(LockerApplication.class, args);
	}

}
