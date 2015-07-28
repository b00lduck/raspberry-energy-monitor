package com.b00lduck.raspberryEnergyMonitor.dataservice;

import com.b00lduck.raspberryEnergyMonitor.dataservice.filter.CorsFilter;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.embedded.FilterRegistrationBean;
import org.springframework.context.annotation.Bean;

/**
 * @author Daniel Zerlett (daniel@zerlett.eu)
 */
@SpringBootApplication
public class Application {

	public static void main(final String[] args) {
		SpringApplication.run(Application.class, args);
	}

	@Bean
	public FilterRegistrationBean commonsRequestLoggingFilter() {
		final FilterRegistrationBean registrationBean = new FilterRegistrationBean();
		registrationBean.setFilter(new CorsFilter());
		return registrationBean;
	}

}
