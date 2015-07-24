package com.b00lduck.raspberryEnergyMonitor.dataservice.controller;

import com.b00lduck.raspberryEnergyMonitor.dataservice.entity.Counter;
import com.b00lduck.raspberryEnergyMonitor.dataservice.entity.CounterEvent;
import com.b00lduck.raspberryEnergyMonitor.dataservice.entity.CounterEventType;
import com.b00lduck.raspberryEnergyMonitor.dataservice.repository.CounterEventRepository;
import com.b00lduck.raspberryEnergyMonitor.dataservice.repository.CounterRepository;
import org.joda.time.DateTime;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.rest.webmvc.RepositoryRestController;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

import javax.transaction.Transactional;
import java.math.BigDecimal;

/**
 * @author Daniel Zerlett (daniel@zerlett.eu)
 * @created 24.07.2015
 */
@RepositoryRestController
public class CounterController {

	@Autowired
	private CounterRepository counterRepository;

	@Autowired
	private CounterEventRepository counterEventRepository;

	@RequestMapping(method = RequestMethod.POST, value = "/counters/{id}/tick")
	public void tickCounter(@PathVariable final Long id) {

		final Counter counter = counterRepository.findOne(id);

		final CounterEvent counterEvent = new CounterEvent();
		counterEvent.setCounter(counter);
		counterEvent.setAbsolute(new BigDecimal("1234567.8943"));
		counterEvent.setType(CounterEventType.COUNT);
		counterEvent.setDelta(new BigDecimal("0.001"));
		counterEvent.setTimestamp(new DateTime());

		counterEventRepository.save(counterEvent);

	}

}
