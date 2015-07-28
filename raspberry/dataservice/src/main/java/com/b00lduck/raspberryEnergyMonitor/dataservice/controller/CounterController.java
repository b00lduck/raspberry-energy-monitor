package com.b00lduck.raspberryEnergyMonitor.dataservice.controller;

import com.b00lduck.raspberryEnergyMonitor.dataservice.entity.Counter;
import com.b00lduck.raspberryEnergyMonitor.dataservice.entity.CounterEvent;
import com.b00lduck.raspberryEnergyMonitor.dataservice.entity.CounterEventType;
import com.b00lduck.raspberryEnergyMonitor.dataservice.repository.CounterEventRepository;
import com.b00lduck.raspberryEnergyMonitor.dataservice.repository.CounterRepository;
import org.joda.time.DateTime;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.rest.webmvc.RepositoryRestController;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import java.math.BigDecimal;

/**
 * @author Daniel Zerlett (daniel@zerlett.eu)
 * created 24.07.2015
 */
@RepositoryRestController
public class CounterController {

	@Autowired
	private CounterRepository counterRepository;

	@Autowired
	private CounterEventRepository counterEventRepository;

	@RequestMapping(method = RequestMethod.POST, value = "/counters/{id}/tick")
	@ResponseBody
	@Transactional
	public ResponseEntity<Counter> tickCounter(@PathVariable final Long id) {

		final Counter counter = counterRepository.findOne(id);

		if (null == counter) {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}

		final BigDecimal delta = new BigDecimal("0.01");

		final BigDecimal newValue = counter.getReading().add(delta);

		final CounterEvent counterEvent = new CounterEvent();
		counterEvent.setCounter(counter);
		counterEvent.setType(CounterEventType.TICK);
		counterEvent.setDelta(delta);
		counterEvent.setReading(newValue);
		counterEvent.setTimestamp(new DateTime());
		counterEventRepository.save(counterEvent);

		counter.setReading(newValue);
		counter.setLastTick(counterEvent.getTimestamp());
		counterRepository.save(counter);

		return new ResponseEntity<>(counter, HttpStatus.OK);

	}

}
