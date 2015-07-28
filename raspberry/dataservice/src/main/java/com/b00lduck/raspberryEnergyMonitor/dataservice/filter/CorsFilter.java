package com.b00lduck.raspberryEnergyMonitor.dataservice.filter;

import org.springframework.stereotype.Component;

import javax.servlet.*;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

/**
 * @author Daniel Zerlett (daniel@zerlett.eu)
 *         created 28.07.2015
 */
@Component
public class CorsFilter implements Filter {

	public void doFilter(final ServletRequest req, final ServletResponse res, final FilterChain chain)
			throws IOException, ServletException {
		final HttpServletResponse response = (HttpServletResponse) res;
		response.setHeader("Access-Control-Allow-Origin", "*");
		response.setHeader("Access-Control-Allow-Methods", "POST, PUT, GET, DELETE");
		response.setHeader("Access-Control-Max-Age", "3600");
		response.setHeader("Access-Control-Allow-Headers", "x-requested-with");
		response.setHeader("Access-Control-Allow-Headers", "Content-Type");
		chain.doFilter(req, res);
	}

	public void init(final FilterConfig filterConfig) {}

	public void destroy() {}

}


