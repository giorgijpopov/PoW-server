package quote

import (
	"context"
	"math/rand"
)

type provider struct{}

var _ Provider = (*provider)(nil)

func NewProvider() Provider {
	return &provider{}
}

func (p *provider) GetQuote(ctx context.Context) (string, error) {
	return quotes[rand.Intn(len(quotes))], nil
}

var quotes = []string{
	"In the end, it's not the years in your life that count. It's the life in your years. Abraham Lincoln",
	"The greatest day in your life and mine is when we take total responsibility for our attitudes. That's the day we truly grow up. John C. Maxwell",
	"Life is not a problem to be solved, but a reality to be experienced. Soren Kierkegaard",
	"What we think determines what happens to us, so if we want to change our lives, we need to stretch our minds. Wayne Dyer",
	"Life is ten percent what happens to you and ninety percent how you respond to it. Lou Holtz",
	"Believe that life is worth living and your belief will help create the fact. William James",
	"The only disability in life is a bad attitude. Scott Hamilton",
	"Too often we underestimate the power of a touch, a smile, a kind word, a listening ear, an honest compliment, or the smallest act of caring, all of which have the potential to turn a life around. Leo Buscaglia",
	"Life isn't about finding yourself. Life is about creating yourself. George Bernard Shaw",
	"There is more to life than increasing its speed. Mahatma Gandhi",
	"Life is really simple, but we insist on making it complicated. Confucius",
	"Our prime purpose in this life is to help others. And if you can't help them, at least don't hurt them. Dalai Lama",
	"There are three constants in life...change, choice and principles. Stephen Covey",
	"Life's most persistent and urgent question is, 'What are you doing for others? Martin Luther King, Jr.",
	"Life is a series of natural and spontaneous changes. Don't resist them--that only creates sorrow. Let reality be reality. Let things flow naturally forward in whatever way they like. Lao Tzu",
	"Change is the law of life. And those who look only to the past or present are certain to miss the future. John F. Kennedy",
	"Only a life lived for others is a life worthwhile. Albert Einstein",
	"When life is too easy for us, we must beware or we may not be ready to meet the blows which sooner or later come to everyone, rich or poor. Eleanor Roosevelt",
	"God gave us the gift of life; it is up to us to give ourselves the gift of living well. Voltaire",
	"We make a living by what we get, but we make a life by what we give. Winston Churchill",
	"All life is an experiment. The more experiments you make the better. Ralph Waldo Emerson",
	"My mission in life is not merely to survive, but to thrive; and to do so with some passion, some compassion, some humor, and some style. Maya Angelou",
	"Once you say you're going to settle for second, that's what happens to you in life. John F. Kennedy",
	"There is no passion to be found playing small--in settling for a life that is less than the one you are capable of living. Nelson Mandela",
	"If you don't design your own life plan, chances are you'll fall into someone else's plan. And guess what they have planned for you? Not much. Jim Rohn",
	"I've failed over and over and over again in my life and that is why I succeed. Michael Jordan",
	"The biggest adventure you can take is to live the life of your dreams. Oprah Winfrey",
	"Literature adds to reality, it does not simply describe it. It enriches the necessary competencies that daily life requires and provides; and in this respect, it irrigates the deserts that our lives have already become. C. S. Lewis",
	"Anyone who stops learning is old, whether at twenty or eighty. Anyone who keeps learning stays young. The greatest thing in life is to keep your mind young. Henry Ford",
	"Many of life's failures are people who did not realize how close they were to success when they gave up. Thomas A. Edison",
	"The most difficult thing is the decision to act, the rest is merely tenacity. The fears are paper tigers. You can do anything you decide to do. You can act to change and control your life; and the procedure, the process is its own reward. Amelia Earhart",
	"People grow through experience if they meet life honestly and courageously. This is how character is built. Eleanor Roosevelt",
	"Remember your dreams and fight for them. You must know what you want from life. There is just one thing that makes your dream become impossible: the fear of failure. Paulo Coelho",
	"Our greatest happiness does not depend on the condition of life in which chance has placed us, but is always the result of a good conscience, good health, occupation, and freedom in all just pursuits. Thomas Jefferson",
	"The quality of a person's life is in direct proportion to their commitment to excellence, regardless of their chosen field of endeavor. Vince Lombardi",
	"Communication is a skill that you can learn. It's like riding a bicycle or typing. If you're willing to work at it, you can rapidly improve the quality of every part of your life. Brian Tracy",
	"Today is life--the only life you are sure of. Make the most of today. Get interested in something. Shake yourself awake. Develop a hobby. Let the winds of enthusiasm sweep through you. Live today with gusto. Dale Carnegie",
	"The secret of success is learning how to use pain and pleasure instead of having pain and pleasure use you. If you do that, you're in control of your life. If you don't, life controls you. Tony Robbins",
	"In three words I can sum up everything I've learned about life: it goes on. Robert Frost",
	"We have always held to the hope, the belief, the conviction that there is a better life, a better world, beyond the horizon. Franklin D. Roosevelt",
	"Life takes on meaning when you become motivated, set goals and charge after them in an unstoppable manner. Les Brown",
	"Life is a daring adventure or nothing at all. Helen Keller",
	"The ultimate value of life depends upon awareness and the power of contemplation rather than upon mere survival. Aristotle",
	"Don't take life too seriously. You'll never get out of it alive. Elbert Hubbard",
	"Each life is made up of mistakes and learning, waiting and growing, practicing patience and being persistent. Billy Graham",
	"Each person must live their life as a model for others. Rosa Parks",
	"My philosophy of life is that if we make up our mind what we are going to make of our lives, then work hard toward that goal, we never lose--somehow we win out. Ronald Reagan",
	"Life is not about how fast you run or how high you climb, but how well you bounce. Vivian Komori",
	"Transformation is a process, and as life happens there are tons of ups and downs. It's a journey of discovery--there are moments on mountaintops and moments in deep valleys of despair. Rick Warren",
	"Live life to the fullest, and focus on the positive. Matt Cameron",
}
