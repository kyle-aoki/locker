package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.CreateEnvironmentDto;
import com.secretsLocker.locker.response.EnvironmentResponse;
import com.secretsLocker.locker.service.EnvironmentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/env")
public class EnvironmentController {

    @Autowired
    EnvironmentService environmentService;

    @PostMapping("/create")
    public EnvironmentResponse.EnvCreated createEnv(
            @RequestBody CreateEnvironmentDto createEnvironmentDto
    ) {
        environmentService.create(createEnvironmentDto);
        return new EnvironmentResponse.EnvCreated();
    }
}
